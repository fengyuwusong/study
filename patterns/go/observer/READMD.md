# 依赖注入，控制反转—观察者模式

- 解决问题：解耦观察者和被观察者，尤其是存在多个观察者的场景。

TCC 使用了观察者模式实现了当某 key 的 value 发生变更时执行回调的逻辑。

TccClient对外提供AddListener方法，允许业务注册对某 key 变更的监听，同时开启定时轮询，如果 key 的值与上次不同就回调业务的 callback 方法。

这里的观察者是调用 AddListener 的发起者，被观察者是 TCC 的 key。Callback可以看作只有一个函数的接口，TccClient的通知回调不依赖于具体的实现，而是依赖于抽象，同时Callback对象不是在内部构建的，而是在运行时传入的，让被观察者不再依赖观察者，通过依赖注入达到控制反转的目的。

参考：https://mp.weixin.qq.com/s/S1BQ55yZgBlB4AkfPX-gIg
