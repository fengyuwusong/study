#include <pthread.h>
#include <stdio.h>
#include <stdatomic.h>
#include <stdlib.h>
#include <time.h>
#define THREAD_COUNT 10
atomic_int accountA = 100000000; // 转出账户初始金额；
atomic_int accountB = 0;         // 转入账户初始金额；
int run(void *v)
{
    int _amount = *((int *)v); // 获得当前线程的转移金额；
    for (;;)
    {
        // 首先判断转出账户金额是否足够，不够则直接退出；
        if (accountA < _amount)
            return thrd_error;
        atomic_fetch_add(&accountB, _amount); // 将金额累加到转入账户；
        atomic_fetch_sub(&accountA, _amount); // 将金额从转出账户中扣除；
    }
}
int main(void)
{
    thrd_t threads[THREAD_COUNT];
    srand(time(NULL));
    for (int i = 0; i < THREAD_COUNT; i++)
    {
        int amount = rand() % 50; // 为每一个线程生成一个随机转移金额；
        thrd_create(&threads[i], run, &amount);
    }
    for (int i = 0; i < THREAD_COUNT; i++)
        thrd_join(threads[i], NULL);
    printf("A: %d\nB: %d", accountA, accountB);
    return 0;
}