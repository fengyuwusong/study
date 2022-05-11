# 替代 IF—策略模式

- 解决场景：支持不同策略的灵活切换，避免多层控制语句的不优雅实现，避免出现如下场景：

```go
if xxx {
  // do something
} else if xxx {
  // do something
} else if xxx {
  // do something
} else if xxx {
  // do something
} else {
}
```

通常的做法是定义了一个公共接口，各种不同的算法以不同的方式实现这个接口，环境角色使用这个接口调用不同的算法。

在 GORM 的 clause/clause.go 中使用到策略模式实现 SQL 的拼装。

现实业务中 SQL 语句千变万化，GORM 将 SQL 的拼接过程，拆分成了一个个小的子句，这些子句统一实现clause.Interface这个接口，然后各自在Build方法中实现自己的构造逻辑。

以最简单的分页查询为例，在使用 db 链式调用构建 SQL 时，对Limit、Offset、Order的函数调用最终转化成了Limit子句和OrderBy子句，两者都实现了clause.Interface接口。

以后 SQL 支持新子句时，创建一个类实现clause.Interface接口，并在函数调用的地方实例化该类，其余执行的代码皆可不变，符合 OOP 中的开闭原则和依赖倒置原则。

参考：https://mp.weixin.qq.com/s/S1BQ55yZgBlB4AkfPX-gIg
