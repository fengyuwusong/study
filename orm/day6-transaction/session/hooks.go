package session

import (
	"geeorm/log"
	"reflect"
)

const (
	BeforeQuery  = "BeforeQuery"
	AfterQuery   = "AfterQuery"
	BeforeUpdate = "BeforeUpdate"
	AfterUpdate  = "AfterUpdate"
	BeforeDelete = "BeforeDelete"
	AfterDelete  = "AfterDelete"
	BeforeInsert = "BeforeInsert"
	AfterInsert  = "AfterInsert"
)

func (s *Session) CallMethod(method string, value interface{}) {
	var fm reflect.Value
	// 如value不为空 则使用value对象获取钩子方法
	if value != nil {
		fm = reflect.ValueOf(value).MethodByName(method)
	} else {
		// 否则获取操作对象的对应操作钩子方法
		fm = reflect.ValueOf(s.RefTable().Model).MethodByName(method)
	}
	// 钩子方法入参均为session
	params := []reflect.Value{reflect.ValueOf(s)}
	if fm.IsValid() {
		if v := fm.Call(params); len(v) > 0 {
			if err, ok := v[0].Interface().(error); ok {
				log.Error(err)
			}
		}
	}
	return
}
