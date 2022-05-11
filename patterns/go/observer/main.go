package main

import (
	"context"
	"fmt"
	"time"
)

// Callback for listener，外部监听者需要实现该方法传入，用于回调
type Callback func(value string, err error)

// 一个监听者实体
type listener struct {
	key             string
	callback        Callback
	lastVersionCode string
	lastValue       string
	lastErr         error
}

// 检测监听的key是否有发生变化，如果有，则回调callback函数
func (l *listener) update(value, versionCode string, err error) {
	if versionCode == l.lastVersionCode && err == l.lastErr {
		return
	}
	if value == l.lastValue && err == l.lastErr {
		// version_code updated, but value not updated
		l.lastVersionCode = versionCode
		return
	}
	defer func() {
		if r := recover(); r != nil {
			logs.Errorf("[TCC] listener callback panic, key: %s, %v", l.key, r)
		}
	}()
	l.callback(value, err)
	l.lastVersionCode = versionCode
	l.lastValue = value
	l.lastErr = err
}

// AddListener add listener of key, if key's value updated, callback will be called
func (c *ClientV2) AddListener(key string, callback Callback, opts ...ListenOption) error {
	listenOps := listenOptions{}
	for _, op := range opts {
		op(&listenOps)
	}

	listener := listener{
		key:      key,
		callback: callback,
	}
	if listenOps.curValue == nil {
		listener.update(c.getWithCache(context.Background(), key))
	} else {
		listener.lastValue = *listenOps.curValue
	}

	c.listenerMu.Lock()
	defer c.listenerMu.Unlock()
	if _, ok := c.listeners[key]; ok {
		return fmt.Errorf("[TCC] listener already exist, key: %s", key)
	}
	c.listeners[key] = &listener
	// 一个client启动一个监听者
	if !c.listening {
		go c.listen()
		c.listening = true
	}
	return nil
}

// 轮询监听
func (c *ClientV2) listen() {
	for {
		time.Sleep(c.listenInterval)
		listeners := c.getListeners()
		for key := range listeners {
			listeners[key].update(c.getWithCache(context.Background(), key))
		}
	}
}

// 加锁防止多线程同时修改listeners，同时拷贝一份map在循环监听时使用。
func (c *ClientV2) getListeners() map[string]*listener {
	c.listenerMu.Lock()
	defer c.listenerMu.Unlock()
	listeners := make(map[string]*listener, len(c.listeners))
	for key := range c.listeners {
		listeners[key] = c.listeners[key]
	}
	return listeners
}
