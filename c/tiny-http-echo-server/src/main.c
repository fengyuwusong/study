#include <sys/socket.h>
#include <netinet/in.h>
#include <pthread.h>
#include <stdatomic.h>
#include <unistd.h>
#include <string.h>
#include <stdlib.h>
#include <stdio.h>
#include <stdnoreturn.h>
#include <signal.h>
#include "../libs/structs.h"
#include "../libs/helpers.h"
#include "../libs/macros.h"

// global variables.
atomic_int threadCounter = 0;
pthread_mutex_t mutex = PTHREAD_MUTEX_INITIALIZER;
pthread_cond_t cond = PTHREAD_COND_INITIALIZER;

void renewThread(void *arg)
{
    // 关闭acceptedSocket文件描述符
    int *acceptedSocket = (int *)arg;
    close(*acceptedSocket);
    pthread_mutex_lock(&mutex);
    threadCounter--;
    pthread_cond_signal(&cond); // 唤醒主线程
    pthread_mutex_unlock(&mutex);
}

noreturn void *acceptConn(void *arg)
{
    acceptParams *ap = (acceptParams *)arg;
    int acceptedSocket;

    while (1)
    {
        // 线程清理回调 压栈方式
        pthread_cleanup_push(renewThread, &acceptedSocket);
        // extracts a request from the queue. 获取acceptedSocket
        if ((acceptedSocket = accept(ap->serverFd, ap->addr, ap->addrLen)) < 0)
        {
            perror("In accept");
            pthread_exit(NULL);
        }

        // deal with HTTP request.
        char reqBuf[HTTP_REQ_BUF];
        bzero(reqBuf, HTTP_REQ_BUF);
        const size_t receivedBytes = read(acceptedSocket, reqBuf, HTTP_REQ_BUF);
        if (receivedBytes > 0)
        {
            char resBuf[HTTP_RES_BUF];

            // retrieve number from query.
            pthread_mutex_lock(&mutex);
            const int num = retrieveGETQueryIntValByKey(reqBuf, "num");
            pthread_mutex_unlock(&mutex);

            int fibResult = calcFibonacci(num);
            // follow the format of the http response.
            sprintf(resBuf, "HTTP/1.1 200 OK\r\n"
                            "Content-type: text/plain\r\n"
                            "Content-length: %d\r\n\r\n%d",
                    calcDigits(fibResult), fibResult);
            write(acceptedSocket, resBuf, strlen(resBuf));
        }
        close(acceptedSocket);
        pthread_cleanup_pop(0);
    }
}

int main(int argc, char const *argv[])
{
    // 初始化配置
    serverSettings ss = {.thread_count = 4};
    setupServerSettings(argc, argv, &ss);

    int serverFd;
    sockaddr_in address;
    int addrLen = sizeof(address);

    // 创建socketFd
    if ((serverFd = socket(AF_INET, SOCK_STREAM, 0)) == 0)
    {
        perror("In socket creation");
        exit(EXIT_FAILURE);
    }

    // 初始化 address
    bzero(&address, addrLen);
    address.sin_family = AF_INET;
    address.sin_addr.s_addr = INADDR_ANY; // -> 0.0.0.0.
    address.sin_port = htons(PORT);

    // 绑定地址到socket
    if (bind(serverFd, (sockaddr *)&address, sizeof(address)) < 0)
    {
        perror("In bind");
        exit(EXIT_FAILURE);
    }

    // 将socket开启监听 变为被动socket
    if (listen(serverFd, MAX_LISTEN_CONN) < 0)
    {
        perror("In listen");
        exit(EXIT_FAILURE);
    }
    printf("\nServer is now listening at port %d:\n\n", PORT);

    // main loop.
    while (1)
    {
        // 当线程数大于最大线程限制时阻塞
        pthread_mutex_lock(&mutex);
        while (threadCounter >= ss.thread_count)
            // 阻塞当前线程 等待其他线程调用方法唤醒
            pthread_cond_wait(&cond, &mutex);
        pthread_mutex_unlock(&mutex);

        // 创建新线程处理当前请求
        pthread_t thread_id;
        acceptParams ap = {serverFd, (sockaddr *)&address, (socklen_t *)&addrLen};
        pthread_create(&thread_id, NULL, acceptConn, &ap);

        // 当前处理线程+1
        atomic_fetch_add(&threadCounter, 1);
        printf("[Info] Thread Created: No.%d\n", threadCounter);
    }
    return EXIT_SUCCESS;
}
