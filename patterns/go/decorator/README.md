# 通过组合扩展功能—装饰模式

- 解决问题：当已有类功能不够便捷时，通过组合的方式实现对已有类的功能扩展，实现了对已有代码的黑盒复用。

TCC 使用了装饰模式扩展了原来已有的ClientV2的能力。

在下面的DemotionClient结构体中组合了ClientV2的引用，对外提供了GetInt和GetBool两个方法，包掉了对原始 string 类型的转换，对外提供了更为便捷的方法。

由于 Golang 语言对嵌入类型的支持，DemotionClient在扩展能力的同时，ClientV2的原本方法也能正常调用，这样语法糖的设计让组合操作达到了继承的效果，且符合 OOP 中替换原则。

参考：https://mp.weixin.qq.com/s/S1BQ55yZgBlB4AkfPX-gIg