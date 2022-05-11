# 参数可选，开箱即用—函数式选项模式

- 解决问题：在设计一个函数时，当存在配置参数较多，同时参数可选时，函数式选项模式是一个很好的选择，它既有为不熟悉的调用者准备好的默认配置，还有为需要定制的调用者提供自由修改配置的能力，且支持未来灵活扩展属性。

这样的方式很好地控制了哪些属性能被外部修改，哪些是不行的。当getoptions需要增加新属性时，给定一个默认值，对应增加一个新GetOption方法即可，对于历史调用方来说无感，能向前兼容式的升级，符合 OOP 中的对修改关闭，对扩展开放的开闭设计原则。

BConfigClient是用于发送 http 请求获取后端服务中 key 对应的 value 值，其中getoptions结构体是 BConfigClient 的配置类，包含请求的 cluster、addr、auth 等信息，小写开头，属于内部结构体，不允许外部直接创建和修改，但同时对外提供了GetOption的方法去修改getoptions中的属性，其中WithCluster、WithAddr、WithAuth是快捷生成GetOption的函数。

NewBConfigClient方法接受一个可变长度的GetOption，意味着调用者可以不用传任何参数，开箱即用，也可以根据自己的需要灵活添加。函数内部首先初始化一个默认配置，然后循环执行GetOption方法，将用户定义的操作赋值给默认配置。

参考：https://mp.weixin.qq.com/s/S1BQ55yZgBlB4AkfPX-gIg