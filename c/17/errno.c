
#include <tgmath.h>
#include <string.h>
#include <stdio.h>
#include <errno.h>
int main(void)
{
    sqrt(-1);
    fprintf(stderr, "%s\n", strerror(errno)); // "Numerical argument out of domain".
    return 0;
}