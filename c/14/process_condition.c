
#include <threads.h>
#include <stdio.h>
mtx_t mutex;
cnd_t cond; // 定义一个条件变量；
int done = 0;
int run(void *data)
{
    mtx_lock(&mutex);
    done = 1;
    cnd_signal(&cond); // 通知等待中的线程；
    mtx_unlock(&mutex);
    return thrd_success;
}
int main(void)
{
#ifndef __STDC_NO_THREADS__
    mtx_init(&mutex, mtx_plain);
    cnd_init(&cond); // 初始化条件变量；
    thrd_t thread;
    thrd_create(&thread, run, NULL);
    mtx_lock(&mutex);
    while (done == 0)
    {
        cnd_wait(&cond, &mutex); // 让当前线程进入等待队列；
    }
    mtx_unlock(&mutex);
    printf("The value of done is: %d", done);
    mtx_destroy(&mutex);
    cnd_destroy(&cond); // 销毁条件变量；
#endif
    return 0;
}