
#include <threads.h>
#include <stdio.h>
#include <stdatomic.h>
#if !defined(__STDC_NO_ATOMICS__)
atomic_int x = 0, y = 0;
#endif
int run(void *v)
{
    atomic_store_explicit(&x, 10, memory_order_relaxed);
    atomic_store_explicit(&y, 20, memory_order_release);
}
int observe(void *v)
{
    while (atomic_load_explicit(&y, memory_order_acquire) != 20)
        ;
    printf("%d", atomic_load_explicit(&x, memory_order_relaxed));
}
int main(void)
{
#if !defined(__STDC_NO_THREADS__) || !defined(__STDC_NO_ATOMICS__)
    thrd_t threadA, threadB;
    thrd_create(&threadA, run, NULL);
    thrd_create(&threadB, observe, NULL);
    thrd_join(threadA, NULL);
    thrd_join(threadB, NULL);
#endif
    return 0;
}