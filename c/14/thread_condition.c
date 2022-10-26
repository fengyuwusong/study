
#include <stdio.h>
#include <threads.h>
#include <stdatomic.h>
#define THREAD_COUNT 10
#define THREAD_LOOP 10000
_Thread_local int counter = 0; // 定义线程本地变量；
int run(void *data)
{
    for (int i = 0; i < THREAD_LOOP; ++i)
        counter += 1; // 更新当前线程所属的 counter 变量值；
    return counter;
}
int main(int argc, char const *argv[])
{
    thrd_t threads[THREAD_COUNT];
    int sum = 0, result = 0;
    for (int i = 0; i < THREAD_COUNT; ++i)
        thrd_create(&threads[i], run, NULL);
    for (int i = 0; i < THREAD_COUNT; ++i)
    {
        thrd_join(threads[i], &result);
        sum += result; // 累加每个线程的计算值；
    }
    printf("The value of count is %d.\n", sum);
    return 0;
}