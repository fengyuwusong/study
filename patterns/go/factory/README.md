# 隐藏复杂对象构造过程—工厂模式

- 解决问题：将对象复杂的构造逻辑隐藏在内部，调用者不用关心细节，同时集中变化。

TCC 创建LogCounnter时使用了工厂模式，该类作用是根据错误日志出现的频率判断是否需要打印日志，如果在指定的时间里，错误日志的触发超过指定次数，则需要记录日志。

NewLogCounter方法通过入参 LogMode 枚举类型即可生成不同规格配置的LogCounnter，可以无需再去理解 TriggerLogCount、TriggerLogDuration、Enable 的含义。

>识别变化隔离变化，简单工厂是一个显而易见的实现方式。它符合了 DRY 原则（Don't Repeat Yourself！），创建逻辑存放在单一的位置，即使它变化，也只需要修改一处就可以了。DRY 很简单，但却是确保我们代码容易维护和复用的关键。DRY 原则同时还提醒我们：对系统职能进行良好的分割，职责清晰的界限一定程度上保证了代码的单一性。[引用自 https://blog.51cto.com/weijie/82767]

参考：https://mp.weixin.qq.com/s/S1BQ55yZgBlB4AkfPX-gIg
