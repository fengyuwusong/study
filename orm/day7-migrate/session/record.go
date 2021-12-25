package session

import (
	"errors"
	"geeorm/clause"
	"reflect"
)

// Insert 将结构体解析平铺成sql语句插入
func (s *Session) Insert(values ...interface{}) (int64, error) {
	recordValues := make([]interface{}, 0)
	for _, value := range values {
		s.CallMethod(BeforeInsert, value)
		table := s.Model(value).RefTable()
		s.clause.Set(clause.INSERT, table.Name, table.FieldNames)
		recordValues = append(recordValues, table.RecordValues(value))
	}

	s.clause.Set(clause.VALUES, recordValues...)
	sql, vars := s.clause.Build(clause.INSERT, clause.VALUES)
	result, err := s.Raw(sql, vars...).Exec()
	if err != nil {
		return 0, err
	}
	s.CallMethod(AfterInsert, nil)
	return result.RowsAffected()
}

// Find 将结构体解析获取每个字段并生成对应sql语句进行查询赋值
func (s *Session) Find(values interface{}) error {
	s.CallMethod(BeforeQuery, nil)
	destSlice := reflect.Indirect(reflect.ValueOf(values))
	// 获取数组的元素类型
	destType := destSlice.Type().Elem()
	// New 元素指正结构体并获取值的interface并解析获得table对象
	table := s.Model(reflect.New(destType).Elem().Interface()).RefTable()

	// 根据表结构信息构造select并查询
	s.clause.Set(clause.SELECT, table.Name, table.FieldNames)
	sql, vars := s.clause.Build(clause.SELECT, clause.WHERE, clause.ORDERBY, clause.LIMIT)
	rows, err := s.Raw(sql, vars...).QueryRows()
	if err != nil {
		return err
	}

	// 获得查询结果并赋值结构体
	for rows.Next() {
		// 实例化单个元素指针
		dest := reflect.New(destType).Elem()
		var values []interface{}
		// 遍历table每个field获取其类型的interface指针
		for _, name := range table.FieldNames {
			values = append(values, dest.FieldByName(name).Addr().Interface())
		}
		// scan将该行数据赋值给构造好的数组指针
		if err := rows.Scan(values...); err != nil {
			return err
		}
		s.CallMethod(AfterQuery, dest.Addr().Interface())
		// 赋值values 循环直至所有元素均添加
		destSlice.Set(reflect.Append(destSlice, dest))
	}
	// 关闭游标
	return rows.Close()
}

func (s *Session) Update(kv ...interface{}) (int64, error) {
	s.CallMethod(BeforeUpdate, nil)
	m, ok := kv[0].(map[string]interface{})
	// 非map 类型则使用平铺key value方式解析构造map
	if !ok {
		m = make(map[string]interface{})
		for i := 0; i < len(kv); i += 2 {
			m[kv[i].(string)] = kv[i+1]
		}
	}
	s.clause.Set(clause.UPDATE, s.RefTable().Name, m)
	sql, vars := s.clause.Build(clause.UPDATE, clause.WHERE)
	result, err := s.Raw(sql, vars...).Exec()
	if err != nil {
		return 0, err
	}
	s.CallMethod(AfterUpdate, nil)
	return result.RowsAffected()
}

func (s *Session) Delete() (int64, error) {
	s.CallMethod(BeforeDelete, nil)
	s.clause.Set(clause.DELETE, s.RefTable().Name)
	sql, vars := s.clause.Build(clause.DELETE, clause.WHERE)
	result, err := s.Raw(sql, vars...).Exec()
	if err != nil {
		return 0, err
	}
	s.CallMethod(AfterDelete, nil)
	return result.RowsAffected()
}

func (s *Session) Count() (int64, error) {
	s.clause.Set(clause.COUNT, s.RefTable().Name)
	sql, vars := s.clause.Build(clause.COUNT, clause.WHERE)
	row := s.Raw(sql, vars...).QueryRow()
	var tmp int64
	if err := row.Scan(&tmp); err != nil {
		return 0, err
	}
	return tmp, nil
}

// Limit adds limit condition to clause
func (s *Session) Limit(num int) *Session {
	s.clause.Set(clause.LIMIT, num)
	return s
}

// Where adds limit condition to clause
func (s *Session) Where(desc string, args ...interface{}) *Session {
	s.clause.Set(clause.WHERE, append([]interface{}{desc}, args...)...)
	return s
}

// OrderBy adds order by condition to clause
func (s *Session) OrderBy(desc string) *Session {
	s.clause.Set(clause.ORDERBY, desc)
	return s
}

func (s *Session) First(value interface{}) error {
	dest := reflect.Indirect(reflect.ValueOf(value))
	// 根据value构造数组且获取对应的value
	destSlice := reflect.New(reflect.SliceOf(dest.Type())).Elem()
	if err := s.Limit(1).Find(destSlice.Addr().Interface()); err != nil {
		return err
	}
	if destSlice.Len() == 0 {
		return errors.New("NOT FOUND")
	}
	dest.Set(destSlice.Index(0))
	return nil
}
