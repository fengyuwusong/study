
#include <stdio.h>
#include <stdlib.h>
#include <time.h>
int main(void)
{
    srand(time(NULL)); // 初始化随机数种子；
    while (getchar() == '\n')
        printf("%d", rand() % 10); // 生成并打印 0-9 的随机数；
    return 0;
}