#include <threads.h>
#include <stdio.h>
#include <stdatomic.h>
#define THREAD_COUNT 10
#define THREAD_LOOP 100000000
#if !defined(__STDC_NO_ATOMICS__)
_Atomic long counter = 0; // 定义一个原子类型全局变量，用来记录线程的累加值；
#endif
int run(void *data)
{
    for (int i = 0; i < THREAD_LOOP; i++)
        atomic_fetch_add_explicit(&counter, 1, memory_order_relaxed); // 使用原子加法操作；
    printf("Thread %d terminates.\n", *((int *)data));
    return thrd_success;
}
int main(void)
{
#if !defined(__STDC_NO_THREADS__) || !defined(__STDC_NO_ATOMICS__)
    int ids[THREAD_COUNT];
    thrd_t threads[THREAD_COUNT];
    for (int i = 0; i < THREAD_COUNT; i++)
    {
        ids[i] = i + 1;
        thrd_create(&threads[i], run, ids + i);
    }
    for (int i = 0; i < THREAD_COUNT; i++)
        thrd_join(threads[i], NULL);
    printf("Counter value is: %ld.\n", counter);
#endif
    return 0;
}