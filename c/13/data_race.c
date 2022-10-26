#include <pthread.h>
#include <stdio.h>
#define THREAD_COUNT 20
#define THREAD_LOOP 100000000
long counter = 0; // 全局变量，用来记录线程的累加值；
void run(void *data)
{
    for (int i = 0; i < THREAD_LOOP; i++)
        counter++; // 在线程中递增全局变量的值；
    printf("Thread %d terminates.\n", *((int *)data));
    return 0;
}
int main(void)
{
    int ids[THREAD_COUNT]; // 用于存放线程序号的数组；
    pthread_t threads[THREAD_COUNT];
    for (int i = 0; i < THREAD_COUNT; i++)
    {
        ids[i] = i + 1;
        pthread_create(&threads[i], NULL, run, ids + i); // 创建 THREAD_COUNT 个线程；
    }
    for (int i = 0; i < THREAD_COUNT; i++)
        pthread_join(threads[i], NULL);          // 让当前线程等待其他线程执行完毕；
    printf("Counter value is: %ld.\n", counter); // 输出 counter 变量最终结果；
    return 0;
}