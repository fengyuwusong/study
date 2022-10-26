
#include <stdio.h>
#include <stdlib.h>
#include <errno.h>
int main(void)
{
    // 一次性字符串到数值转换；
    const char *strA = "1.0";
    printf("%f\n", atof(strA));
    // 带溢出检查的转换函数，执行后会保存不能被转换部分的地址；
    const char *strB = "200000000000000000000000000000.0";
    char *end;
    double num = strtol(strB, &end, 10);
    if (errno == ERANGE)
    { // 判断转换结果是否发生溢出；
        printf("Range error, got: ");
        errno = 0;
    }
    printf("%f\n", num);
    return 0;
}