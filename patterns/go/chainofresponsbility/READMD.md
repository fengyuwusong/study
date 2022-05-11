# Web 中间件—责任链模式

- 解决问题：当业务处理流程很长时，可将所有请求的处理者通过前一对象记住其下一个对象的引用而连成一条链；当有请求发生时，可将请求沿着这条链传递，直到没有对象处理它为止。

APIX 应用了责任链模式来实现中间件的功能，类似的逻辑可参考文章“Gin 中间件的编写和使用”。

首先要定义中间件接口，即下文中的HandlerFunc，然后定义HandlersChain将一组处理函数组合成一个处理链条，最后将HandlersChain插入Context中。

开始执行时，是调用Context的Next函数，遍历每个HandlerFunc，然后将Context自身的引用传入，index是记录当前执行到第几个中间件，当过程中出现不满足继续进行的条件时，可以调用Abort()来终止流程。

参考：https://mp.weixin.qq.com/s/S1BQ55yZgBlB4AkfPX-gIg
