
#include <stdio.h>
#include <time.h>
int main(void)
{
    time_t currTime = time(NULL);
    if (currTime != (time_t)(-1)){
        printf("The current timestamp is: %ld(s)\n", currTime);
        char buff[64];
        struct tm* tm = localtime(&currTime);
        if (strftime(buff, sizeof(buff), "%A %c", tm))
            printf("The current local time is: %s\n", buff);
    }
    return 0;
}