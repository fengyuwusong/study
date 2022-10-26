
#include <signal.h>
#include <stdio.h>
int main(void)
{
    signal(SIGTERM, SIG_IGN); // 忽略信号 SIGTERM；
    raise(SIGTERM);           // 向当前程序发送 SIGTERM 信号；
    printf("Reachable!\n");   // Reachable code!
    return 0;
}