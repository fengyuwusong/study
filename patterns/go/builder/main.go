package main

import "strings"

// Conditional handlers chain
type CondHandlersChain struct {
	// 匹配函数
	Matcher func(method, path string) bool
	// 命中匹配后，执行的处理函数
	Chain HandlersChain
}

func main() {
	// 对路径前缀为 `/wechat` 的请求开启微信认证中间件
	mw1 := apix.CondHandlersChain{
		// SimpleMatcherBuilder是一个建造者, 隐藏细节, 设置前缀等信息后直接build即可
		Matcher: new(apix.SimpleMatcherBuilder).PrefixPath("/wechat").Build(),
		Chain:   apix.HandlersChain{wxsession.NewMiddleware()},
	}

	// 注册中间件
	e.CondUse(mw1)
}

type MatcherBuilder interface {
	Build() func(method, path string) bool
}

var _ MatcherBuilder = (*SimpleMatcherBuilder)(nil)

// SimpleMatcherBuilder build a matcher for CondHandlersChain.
// An `AND` logic will be applied to all fields(if provided).
type SimpleMatcherBuilder struct {
	method     string
	pathPrefix string
	paths      []string
}

func (m *SimpleMatcherBuilder) Method(method string) *SimpleMatcherBuilder {
	m.method = method
	return m
}

func (m *SimpleMatcherBuilder) PrefixPath(path string) *SimpleMatcherBuilder {
	m.pathPrefix = path
	return m
}

func (m *SimpleMatcherBuilder) FullPath(path ...string) *SimpleMatcherBuilder {
	m.paths = append(m.paths, path...)
	return m
}

func (m *SimpleMatcherBuilder) Build() func(method, path string) bool {
	method, prefix := m.method, m.pathPrefix
	paths := make(map[string]struct{}, len(m.paths))
	for _, p := range m.paths {
		paths[p] = struct{}{}
	}

	return func(m, p string) bool {
		if method != "" && m != method {
			return false
		}
		if prefix != "" && !strings.HasPrefix(p, prefix) {
			return false
		}

		if len(paths) == 0 {
			return true
		}

		_, ok := paths[p]
		return ok
	}
}

var _ MatcherBuilder = (AndMBuilder)(nil)
var _ MatcherBuilder = (OrMBuilder)(nil)
var _ MatcherBuilder = (*NotMBuilder)(nil)
var _ MatcherBuilder = (*ExcludePathBuilder)(nil)
