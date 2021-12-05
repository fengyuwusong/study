package gee

import (
	"fmt"
	"strings"
)

// 路由前缀树实现 FIXME 路由前缀树仍存在覆盖 冲突问题, 后续需继续解决

type node struct {
	pattern  string  // 待匹配路由 例如: /p/:lang
	part     string  // 当前节点路由中一部分 例如: :lang
	children []*node // 子节点
	isWild   bool    // 是否精确匹配
}

// 第一个匹配成功的节点 用于插入
func (n *node) matchChild(part string) *node {
	for _, child := range n.children {
		if child.part == part || child.isWild {
			return child
		}
	}
	return nil
}

// 所有匹配成功的节点 用于查找
func (n *node) matchChildren(part string) []*node {
	nodes := make([]*node, 0)
	for _, child := range n.children {
		if child.part == part || child.isWild {
			nodes = append(nodes, child)
		}
	}
	return nodes
}

// 前缀树插入
func (n *node) insert(pattern string, parts []string, height int) {
	// len(parts) == height 意为当前为最后一层 递归出口 赋值pattern绑定即可
	if len(parts) == height {
		n.pattern = pattern
		return
	}

	// 获取该层的part
	part := parts[height]
	// 获取第一个子节点
	child := n.matchChild(part)
	// 不存在则新建
	if child == nil {
		child = &node{
			part:   part,
			isWild: part[0] == ':' || part[0] == '*',
		}
		n.children = append(n.children, child)
	}

	// 递归后续节点
	child.insert(pattern, parts, height+1)
}

// 前缀树搜索 获取子节点
func (n *node) search(parts []string, height int) *node {
	// 递归出口 到达对应层级
	if len(parts) == height || strings.HasPrefix(n.part, "*") {
		if n.pattern == "" {
			return nil
		}
		return n
	}

	// 获取该层级对应的part
	part := parts[height]
	// 获取所有匹配成功的节点
	children := n.matchChildren(part)
	// 遍历节点 递归搜索
	for _, child := range children {
		result := child.search(parts, height+1)
		if result != nil {
			return result
		}
	}
	return nil
}

func (n *node) travel(list *[]*node) {
	if n.pattern != "" {
		*list = append(*list, n)
	}
	for _, child := range n.children {
		child.travel(list)
	}
}

func (n *node) String() string {
	return fmt.Sprintf("node{pattern=%s, part=%s, isWild=%t}", n.pattern, n.part, n.isWild)
}
