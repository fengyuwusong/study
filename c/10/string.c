
#include <string.h>
#include <stdio.h>
#include <limits.h>
#include <ctype.h>

#define LEN 128

int main(void)
{
    // 字符定义方式
    char a = 'c';

    // 字符类型占用位数
    printf("char type has %lu byte.\n", sizeof(char));
    printf("char type has %d bits.\n", CHAR_BIT);

    // 字符串定义方式
    const char token[] = "Hello, ";
    const char *token2 = "geek!";
    printf("%s\n", token);

    // strtok
    char str1[] = "test";
    char* str2 = "t";
    printf("%s\n", strtok(str1, str2));

    // 获取长度
    printf("字符串长度: %zu\n", strlen(token));

    // 拼接
    printf("字符串拼接: %s\n", strncat(token, token2, strlen(token2)));

    // 拷贝字符串
    const char token3[] = "geek!";
    printf("拷贝字符串: %s\n", strncpy(token3, token, strlen(token3)));
    printf("拷贝字符串: %s\n", strncpy(token3, token, sizeof(token3)));

    // 格式化字符串
    char dest[LEN];
    const char strA[] = "Hello, ";
    sprintf(dest, "%sword!", strA);
    printf("%s\n", dest);

    // 字符判断和转换
    char c = 'a';
    printf("%d\n", isalnum(c)); // 8.
    printf("%d\n", isalpha(c)); // 1024.
    printf("%d\n", isblank(c)); // 0.
    printf("%d\n", isdigit(c)); // 0.
    printf("%c\n", toupper(c)); // 'A'.

    return 0;
}