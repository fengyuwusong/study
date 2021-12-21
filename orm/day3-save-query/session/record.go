package session

import (
	"geeorm/clause"
	"reflect"
)

// Insert 将结构体解析平铺成sql语句插入
func (s *Session) Insert(values ...interface{}) (int64, error) {
	recordValues := make([]interface{}, 0)
	for _, value := range values {
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

	return result.RowsAffected()
}

// Find 将结构体解析获取每个字段并生成对应sql语句进行查询赋值
func (s *Session) Find(values interface{}) error {
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
		// 遍历table每个field获取其类型的interface指正
		for _, name := range table.FieldNames {
			values = append(values, dest.FieldByName(name).Addr().Interface())
		}
		// scan将该行数据赋值给构造好的数组指针
		if err := rows.Scan(values...); err != nil {
			return err
		}
		// 赋值values 循环直至所有元素均添加
		destSlice.Set(reflect.Append(destSlice, dest))
	}
	// 关闭游标
	return rows.Close()
}
