
#include <stdio.h>
// restrict 使之不会发生aliasing:内存中的某一个位置，可以通过程序中多于一个的变量来访问或修改其包含的数据。
// 而这可能会导致一个潜在的问题：即当通过其中的某个变量修改数据时，便会导致所有与其他变量相关的数据访问发生改变
// 仅读取一次z指针值
void foo(int *x, int *y, int *restrict z)
{
    *x += *z;
    *y += *z;
}
int main(void)
{
    int x = 10, y = 20, z = 30;
    foo(&x, &y, &z);
    printf("%d %d %d", x, y, z);
    return 0;
}