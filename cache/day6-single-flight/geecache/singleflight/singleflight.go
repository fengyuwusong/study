package singleflight

import "sync"

// 进行中的请求
type call struct {
	wg  sync.WaitGroup
	val interface{}
	err error
}

// 管理不同key的请求
type Group struct {
	mu sync.Mutex
	m  map[string]*call
}

func (g *Group) Do(key string, fn func() (interface{}, error)) (interface{}, error) {
	g.mu.Lock()
	if g.m == nil {
		g.m = make(map[string]*call)
	}

	// 获取当前请求的call
	if c, ok := g.m[key]; ok {
		g.mu.Unlock()       // map操作结束 解锁
		c.wg.Wait()         // 如果请求正在进行 则等待
		return c.val, c.err // 请求结束 返回结果
	}

	c := new(call)
	c.wg.Add(1)   // 发起请求前加锁
	g.m[key] = c  // 添加到 g.m 表明 key 已有对应请求正在处理
	g.mu.Unlock() // map操作结束 解锁

	c.val, c.err = fn() // 调用fn 发起请求
	c.wg.Done()         // 请求结束

	g.mu.Lock()
	delete(g.m, key) // 由于当次请求结束 删除map对应call
	g.mu.Unlock()

	return c.val, c.err // 返回结果
}
