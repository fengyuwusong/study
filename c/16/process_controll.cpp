
#include <stdio.h>
#include <stdlib.h>
void exitHandler()
{
    printf("%s\n", getenv("PATH"));
}
int main(void)
{
    // 注册退出回调函数
    if (!atexit(exitHandler))
    {
        // 与宿主机命令行通信
        system("ls");
        exit(EXIT_SUCCESS);
    }
    return 0;
}