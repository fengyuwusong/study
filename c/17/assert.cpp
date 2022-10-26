#include <stdio.h>
#include <stdlib.h>
#include <assert.h>
double sqrt(double x)
{
    // 检查函数使用时传入的参数；
    assert(x > 0.0);
    // ...
}
int main(void)
{
    // 检查程序的编译要求；
    static_assert(sizeof(int) >= 4,
                  "Integer should have at least 4 bytes length.");
    // ...
    return 0;
}