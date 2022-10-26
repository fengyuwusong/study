#include <stdio.h>
#include <pthread.h>

void *fun1()
{                        // 没有参数
    puts("fun1() end."); // 显示一些信息，便于判断多个线程的结束顺序
    return 0;
}

typedef struct
{
    int age;
    char name[10];
} Horse;

void *fun2(void *p)
{ // 用一个结构体传入参数
    Horse h = *(Horse *)p;
    printf("age: %d, name: %s.\n", h.age, h.name);
    puts("fun2() end.");
    return 0;
}

int main()
{
    int rc;
    pthread_t id;

    // 创建第一个线程
    rc = pthread_create(&id, NULL, fun1, NULL);
    if (rc)
        puts("Failed to create the thread fun1().");

    // 创建第二个线程
    Horse horse = {5, "Jack"};
    rc = pthread_create(&id, NULL, fun2, &horse);
    if (rc)
        puts("Failed to create the thread fun2().");

    // 阻塞主线程的运行，以免主线程运行结束时提前终止子线程
    pthread_join(id, NULL);
    puts("main() end.");
    return 0;
}
