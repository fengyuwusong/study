
#include <stdio.h>
#include <stdarg.h>
void print_sum(int count, ...)
{
    int sum = 0;
    va_list ap;
    va_start(ap, count);
    for (int i = 0; i < count; ++i)
        sum += va_arg(ap, int);
    va_end(ap);
    printf("%d\n", sum);
}
int main(void)
{
    print_sum(4, 1, 2, 3, 4);
    return 0;
}