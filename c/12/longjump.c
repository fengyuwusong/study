#include <stdio.h>
#include <setjmp.h>
#include <stdnoreturn.h>

jmp_buf jb;

noreturn void insepct(char val)
{
    putchar(val);
    longjmp(jb, val);
}

int main(int argc, char const *argv[])
{
    volatile char c = 'A';
    if (setjmp(jb) < 'J')
    {
        insepct(c++);
    }
    return 0;
}
