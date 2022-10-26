#include <stdio.h>

#define ARR_LEN 5
int main(void)
{
    int arr[ARR_LEN] = {1, 5, 10, 9, 0};
    for (int i = 0; i < ARR_LEN; ++i)
    {
        if (arr[i] > 7)
        {
            // save this element somewhere else.
            printf("%d\n", arr[i]);
        }
    }
    return 0;
}