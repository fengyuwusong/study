
#include <signal.h>
#include <stdio.h>
#include <unistd.h>
#include <stdlib.h>
// 可重入变量 信号中断方法执行对该变量读写是原子的
volatile sig_atomic_t sig = 0;
void sigHandler(int signal)
{
    sig = signal;
    printf("sigHandler!!\n");
}
int main(void)
{
    signal(SIGINT, sigHandler);
    int counter = 0; // 计数器变量；
    while (1)
    {
        switch (sig)
        { // 信号筛选与处理；
        case SIGINT:
        {
            printf("SignalValue: %d", sig);
            /* 异常处理的主要逻辑 */
            exit(SIGINT);
        }
        }
        if (counter == 5)
            raise(SIGINT);
        printf("Counter: %d\n", counter++);
        sleep(1);
    }
    return 0;
}