
#include <stdio.h>
int main(void)
{
    printf("Enter some characters:\n");
    FILE *fp = fopen("./temp.txt", "w+");
    // 设置缓存区
    char buf[1024];
    setvbuf(fp, buf, _IOFBF, 5);
    
    if (fp)
    {
        char ch;
        while (scanf("%c", &ch))
        {
            if (ch == 'z')
                break;
            putc(ch, fp);
        }
    }
    else
    {
        perror("File open failed.");
    }
    fclose(fp);
    return 0;
}