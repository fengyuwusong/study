package main

import (
	"fmt"
	"reflect"
)

type parent struct {
	monName    string
	fatherName string
}

type class struct {
	name string
}

type student struct {
	id     int
	ext    map[string]interface{}
	parent *parent
	class  class
	source []int
}

func main() {
	a := &student{
		id: 0,
		ext: map[string]interface{}{
			"test": "test",
		},
		parent: &parent{
			monName:    "a_mon",
			fatherName: "a_father",
		},
		class: class{
			name: "a_class",
		},
		source: []int{1},
	}

	b := a.clone()

	fmt.Println("clone方法相当于开辟了一个新的空间用于存储b，并返回对应指正 故两指针地址比较为false")
	// false clone方法相当于开辟了一个新的空间用于存储b，并返回对应指正 故两指针地址比较为false
	fmt.Printf("a == b: %v\n", a == b)
	fmt.Println("reflect.DeepEqual 方法当两指正地址不同时，会深度比较两个指针的值 值相等故为true")
	// true reflect.DeepEqual 方法当两指正地址不同时，会深度比较两个指针的值 值相等故为true
	fmt.Printf("reflect.DeepEqual a b, result: %v\n", reflect.DeepEqual(a, b))

	fmt.Println("值成员 存储地址不同 故修改不同步")
	// 值成员 存储地址不同 故修改不同步 false
	a.id = 2
	fmt.Printf("a.id == b.id: %v\n", a.id == b.id)
	b.class.name = "b_class"
	fmt.Printf("b.class.name == a.class.name: %v\n", a.class.name == b.class.name)

	// 指针成员引用地址一致 故修改同步 true
	fmt.Println("指针成员引用地址一致 故修改同步")
	a.ext["test"] = "test_a"
	a.ext["new_field"] = "test_a"
	fmt.Printf("a.ext[\"test\"] == b.ext[\"test\"]: %v\n", a.ext["test"] == b.ext["test"])
	fmt.Printf("a.ext[\"new_field\"] == b.ext[\"new_field\"]: %v\n", a.ext["new_field"] == b.ext["new_field"])
	b.parent.monName = "b_mon"
	fmt.Printf("a.parent.monName == b.parent.monName: %v\n", a.parent.monName == b.parent.monName)
	b.source[0] = 2
	fmt.Printf("a.source[0] == b. source[0]: %v\n", a.source[0] == b.source[0])
}

func (s *student) clone() *student {
	copyStudent := *s
	return &copyStudent
}
