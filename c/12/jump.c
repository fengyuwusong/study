#include <stdio.h>

int main(int argc, char const *argv[])
{
    int n;
head:
    scanf("%d", &n);
    if (n < 0) {
        goto head;
    }
    return 0;
}
