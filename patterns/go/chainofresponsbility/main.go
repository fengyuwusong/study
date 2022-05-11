package main

// 定义中间件的接口
type HandlerFunc func(*Context)

// 将一组处理函数组合成一个处理链条
type HandlersChain []HandlerFunc

// 处理的上下文
type Context struct {
	// ...

	// handlers 是一个包含执行函数的数组
	// type HandlersChain []HandlerFunc
	handlers HandlersChain
	// index 表示当前执行到哪个位置了
	index int8

	// ...
}

// Next 会按照顺序将一个个中间件执行完毕
// 并且 Next 也可以在中间件中进行调用，达到请求前以及请求后的处理
func (c *Context) Next() {
	c.index++
	for c.index < int8(len(c.handlers)) {
		if handler := c.handlers[c.index]; handler != nil {
			handler(c)
		}
		c.index++
	}
}

// 停止中间件的循环, 通过将索引后移到abortIndex实现。
func (c *Context) Abort() {
	if c.IsDebugging() && c.index < int8(len(c.handlers)) {
		handler := c.handlers[c.index]
		handlerName := nameOfFunction(handler)
		c.SetHeader("X-APIX-Aborted", handlerName)
	}

	c.index = abortIndex
}
