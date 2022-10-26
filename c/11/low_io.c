
#include <unistd.h>
#include <fcntl.h>
int main(void)
{
    const char str[] = "Enter some characters:\n";
    write(STDOUT_FILENO, str, sizeof(str));
    const int fd = open("./temp.txt", O_RDWR | O_CREAT);
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