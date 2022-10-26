
#include <stdio.h>
#include <time.h>
#include <stdlib.h>
#include <threads.h>
typedef struct timespec ts_t;
int run(void *data)
{ // 模拟的耗时任务；
    volatile double d = 0;
    for (int n = 0; n < 10000; ++n)
        for (int m = 0; m < 10000; ++m)
            d += d * n * m;
    return 0;
}
int main(void)
{
    // 首次记录日历时间与处理器时间；
    ts_t ts1;
    timespec_get(&ts1, TIME_UTC);
    clock_t t1 = clock();
    // 创建两个线程，做一些耗时任务；
    thrd_t thr1, thr2;
    thrd_create(&thr1, run, NULL);
    thrd_create(&thr2, run, NULL);
    thrd_join(thr1, NULL);
    thrd_join(thr2, NULL);
    // 再次记录日历时间与处理器时间；
    ts_t ts2;
    timespec_get(&ts2, TIME_UTC);
    clock_t t2 = clock();
    // 分别计算和打印处理器时间与墙上时钟时间耗时；
    printf("CPU time used (per clock()): %.2f ms\n", 1000.0 * (t2 - t1) / CLOCKS_PER_SEC);
    printf("Wall time passed: %.2f ms\n",
           1000.0 * ts2.tv_sec + 1e-6 * ts2.tv_nsec - (1000.0 * ts1.tv_sec + 1e-6 * ts1.tv_nsec));
    return 0;
}