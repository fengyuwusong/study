
#include <time.h>
#include <stdio.h>
int main(void)
{
    clock_t startTime = clock();
    for (int i = 0; i < 10000000; i++)
    {
    }
    clock_t endTime = clock();
    printf("Consumed CPU time is：%fs\n",
           (double)(endTime - startTime) / CLOCKS_PER_SEC);
    return 0;
}