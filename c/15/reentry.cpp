
#include <stdio.h>
#include <signal.h>
#include <string.h>
#define BUF_SIZE 16 // 全局静态数组大小；
#define FORMAT_NUM_(N) " $" #N
#define FORMAT_NUM(N) FORMAT_NUM_(N)
#define RAISE_EXP_false_ASM()
// 调用 raise 函数向当前程序发送信号；
#define RAISE_EXP_true_ASM() \
    "movl    $4, %%edi\n\t"  \
    "call    raise\n\t"
// 内联汇编实现；
#define INLINE_ASM(ID, HAS_EXP)                                                      \
    "mov     %0, %%r8\n\t" /* 复制传入的字符串数据到全局静态数组 */ \
    "testq   %%rsi, %%rsi\n\t"                                                       \
    "je      .L1" #ID "\n\t"                                                         \
    "xorl    %%eax, %%eax\n\t"                                                       \
    ".L3" #ID ":\n\t"                                                                \
    "movzbl  (%%rdi,%%rax), %%ecx\n\t"                                               \
    "movb    %%cl, (%%r8,%%rax)\n\t"                                                 \
    "addq    $1, %%rax\n\t"                                                          \
    "cmpq    %%rsi, %%rax\n\t"                                                       \
    "jne     .L3" #ID "\n\t"                                                         \
    ".L1" #ID ":\n\t" RAISE_EXP_##HAS_EXP##_ASM() /* 选择性调用 raise 函数 */ \
        "mov     $1, %%rax\n\t"                                                      \
        "mov     $1, %%rdi\n\t"                                                      \
        "mov     %0, %%rsi\n\t"                                                      \
        "mov" FORMAT_NUM(BUF_SIZE) ", %%rdx\n\t"                                     \
                                   "syscall\n\t" /* 触发系统调用，打印内容 */

static char buf[BUF_SIZE]; // 用于保存字符的全局静态数组；
void print_with_exp(const char *str, size_t len)
{ // 会引起信号中断的版本；
    asm(INLINE_ASM(a, true)::"g"(buf));
}
void print_normal(const char *str, size_t len)
{ // 正常的版本；
    asm(INLINE_ASM(b, false)::"g"(buf));
}
void sigHandler(int sig)
{
    const char *str = "Hello";
    print_normal(str, strlen(str));
}
int main(void)
{
    signal(SIGILL, sigHandler);
    const char *str = ", world!";
    print_with_exp(str, strlen(str));
    return 0;
}