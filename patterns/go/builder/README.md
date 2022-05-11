# 一步步构建复杂对象—建造者模式

- 解决问题：使用多个简单的对象一步一步构建成一个复杂的对象。

APIX 在创建请求的匹配函数Matcher时使用了建造者模式。

APIX 中提供了指定对哪些 request 生效的中间件，定义和使用方式如下，CondHandlersChain结构体中定义了匹配函数Matcher和命中后执行的处理函数HandlersChain。

以“对路径前缀为`/wechat` 的请求开启微信认证中间件”为例子，Matcher 函数不用开发者从头实现一个，只需要初始化 SimpleMatcherBuilder 对象，设置请求前缀后，直接 Build 出来即可，它将复杂的匹配逻辑隐藏在内部，非常好用。

SimpleMatcherBuilder是一个建造者，它实现了MatcherBuilder接口，该类支持 method、pathPrefix 和 paths 三种匹配方式，业务方通过Method()、PrefixPath()、FullPath()三个方法的组合调用即可构造出期望的匹配函数。

除此之外，ExcludePathBuilder，AndMBuilder、OrMBuilder、*NotMBuilder也实现了MatcherBuilder接口，某些对象内部又嵌套了对MatcherBuilder的调用，达到了多条件组合起来匹配的目的，非常灵活。

参考：https://mp.weixin.qq.com/s/S1BQ55yZgBlB4AkfPX-gIg
