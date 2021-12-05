package gee

import (
	"net/http"
	"strings"
)

type router struct {
	roots    map[string]*node       // 用于区分不同方法的前缀树 存储动态路由
	handlers map[string]HandlerFunc // 用于存储静态路由
}

func newRouter() *router {
	return &router{
		roots:    make(map[string]*node),
		handlers: make(map[string]HandlerFunc),
	}
}

// 解析动态路由为parts 当存在*时，后续将不在添加解析
func parsePattern(pattern string) []string {
	vs := strings.Split(pattern, "/")

	parts := make([]string, 0)
	for _, item := range vs {
		if item != "" {
			parts = append(parts, item)
			if item[0] == '*' {
				break
			}
		}
	}
	return parts
}

func (r *router) addRoute(method string, pattern string, handler HandlerFunc) {
	parts := parsePattern(pattern)

	key := method + "-" + pattern
	// 初始化方法根节点
	if _, ok := r.roots[method]; !ok {
		r.roots[method] = &node{}
	}
	// 插入对应方法动态根节点
	r.roots[method].insert(pattern, parts, 0)
	// 写入静态路由
	r.handlers[key] = handler
}

// 解析路由 返回对应子节点以及参数
func (r *router) getRoute(method string, path string) (*node, map[string]string) {
	searchParts := parsePattern(path)
	params := make(map[string]string)
	root, ok := r.roots[method]

	// 注册路由无该方法
	if !ok {
		return nil, nil
	}

	// 获取动态路由子节点
	n := root.search(searchParts, 0)
	// 子节点不为空时 解析参数
	if n != nil {
		parts := parsePattern(n.pattern)
		for index, part := range parts {
			if part[0] == ':' {
				// key移除':' 存储map
				params[part[1:]] = searchParts[index]
			}
			if part[0] == '*' && len(part) > 1 {
				// key移除'*' value为searchParts index及其后续的所有part 使用'/'组装
				params[part[1:]] = strings.Join(searchParts[index:], "/")
				break
			}
		}
		return n, params
	}
	return nil, nil
}

func (r *router) handle(c *Context) {
	// 获取对应node
	n, params := r.getRoute(c.Method, c.Path)
	if n != nil {
		c.Params = params
		key := c.Method + "-" + n.pattern
		c.handlers = append(c.handlers, r.handlers[key])
	} else {
		c.handlers = append(c.handlers, func(c *Context) {
			c.String(http.StatusNotFound, "404 NOT FOUND: %s\n", c.Path)
		})
	}
	c.Next()
}
