
#include <stdio.h>
#include <signal.h>
#include <stdlib.h>
void sigHandler(int sig)
{
    printf("Signal %d catched!\n", sig);
    exit(sig);
}
int main(void)
{
    signal(SIGFPE, sigHandler);
    int x = 10;
    int y = 0;
    printf("%d", x / y);
}