package main

import "fmt"

type Getter interface {
	Get(key string) string
}

type GetterFunc func(key string) string

func (f GetterFunc) Get(key string) string {
	// 此处可单独进行其他操作 相当于代理模式
	return f(key)
}

// Get 实际调用方法 接收Getter进行方法调用 避免传入匿名函数方式不够优雅
func Get(getter Getter, key string) string {
	return getter.Get(key)
}

func Get2(f GetterFunc, key string) string {
	return f(key)
}

// 接口型函数总结：
// 1. 相对于原始匿名函数实现，使用接口型函数写法更加优雅易读
// 2. 相对于直接传入函数类型，使用接口型函数可在抽象中实现代理模式功能
// 3. 可在封装上选择性更多，例如抽象后不仅可使用函数实现，也可以使用struct实现，抽象和复用程度更高
func main() {
	f := GetterFunc(func(key string) string {
		return fmt.Sprintf("get1, key: %s", key)
	})

	println(Get(f, "test"))

	f2 := func(key string) string {
		return fmt.Sprintf("get2, key: %s", key)
	}
	println(Get2(f2, "test"))
}
