
#include <unistd.h>
#include <fcntl.h>
int main(void)
{
    const char str[] = "Enter some characters:\n";
    write(STDOUT_FILENO, str, sizeof(str));
    const char *fileName = "./temp.txt";
    // Call to `open` starts:
    // const int fd = open("./temp.txt", O_RDWR | O_CREAT);
    volatile int fd;
    asm("mov $2, %%rax\n\t"
        "mov %0, %%rdi\n\t"
        "mov $66, %%rsi\n\t" // 2 | 64 -> 66;
        "syscall\n\t"
        "mov %%rax, %1\n\t"
        : "=m"(fileName)
        : "m"(fd));
    // Call ended.
    if (fd > 0)
    {
        char ch;
        while (read(STDIN_FILENO, &ch, 1))
        {
            if (ch == 'z')
                break;
            write(fd, &ch, sizeof(ch));
        }
    }
    else
    {
        const char errMsg[] = "File open failed.";
        write(STDERR_FILENO, errMsg, sizeof(errMsg));
    }
    close(fd);
    return 0;
}