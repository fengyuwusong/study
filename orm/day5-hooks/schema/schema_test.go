package schema

import (
	"geeorm/dialect"
	"reflect"
	"testing"
)

type User struct {
	Name string `geeorm:"PRIMARY KEY"`
	Age  int
}

var TestDial, _ = dialect.GetDialect("sqlite3")

func TestParse(t *testing.T) {
	type args struct {
		dest interface{}
		d    dialect.Dialect
	}
	var tests = []struct {
		name string
		args args
		want *Schema
	}{{
		"base",
		args{
			dest: &User{},
			d:    TestDial,
		},
		&Schema{
			Model: &User{},
			Name:  "User",
			Fields: []*Field{{
				Name: "Name",
				Type: "text",
				Tag:  "PRIMARY KEY",
			}, {
				Name: "Age",
				Type: "integer",
				Tag:  "",
			}},
			FieldNames: []string{"Name", "Age"},
			fieldMap: map[string]*Field{
				"Name": {
					Name: "Name",
					Type: "text",
					Tag:  "PRIMARY KEY",
				},
				"Age": {
					Name: "Age",
					Type: "integer",
					Tag:  "",
				},
			},
		},
	}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Parse(tt.args.dest, tt.args.d); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Parse() = %v, want %v", got, tt.want)
			}
		})
	}
}
