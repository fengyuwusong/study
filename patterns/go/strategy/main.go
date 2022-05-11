package main

import "strconv"

func main() {
	db.WithContext(ctx).
		Model(&Course{}).
		Order("course_id DESC").
		Limit(0).
		Offset(100)
}

// Interface clause interface
type Interface interface {
	Name() string
	Build(Builder)
	MergeClause(*Clause)
}

// Limit limit clause
type Limit struct {
	Limit  int
	Offset int
}

// Build build where clause
func (limit Limit) Build(builder Builder) {
	if limit.Limit > 0 {
		builder.WriteString("LIMIT ")
		builder.WriteString(strconv.Itoa(limit.Limit))
	}
	if limit.Offset > 0 {
		if limit.Limit > 0 {
			builder.WriteString(" ")
		}
		builder.WriteString("OFFSET ")
		builder.WriteString(strconv.Itoa(limit.Offset))
	}
}

// Limit specify the number of records to be retrieved
func (db *DB) Limit(limit int) (tx *DB) {
	tx = db.getInstance()
	tx.Statement.AddClause(clause.Limit{Limit: limit})
	return
}

// Offset specify the number of records to skip before starting to return the records
func (db *DB) Offset(offset int) (tx *DB) {
	tx = db.getInstance()
	tx.Statement.AddClause(clause.Limit{Offset: offset})
	return
}

// Order specify order when retrieve records from database
//     db.Order("name DESC")
//     db.Order(clause.OrderByColumn{Column: clause.Column{Name: "name"}, Desc: true})
func (db *DB) Order(value interface{}) (tx *DB) {
	tx = db.getInstance()

	switch v := value.(type) {
	case clause.OrderByColumn:
		tx.Statement.AddClause(clause.OrderBy{
			Columns: []clause.OrderByColumn{v},
		})
	case string:
		if v != "" {
			tx.Statement.AddClause(clause.OrderBy{
				Columns: []clause.OrderByColumn{{
					Column: clause.Column{Name: v, Raw: true},
				}},
			})
		}
	}
	return
}

type OrderByColumn struct {
	Column  Column
	Desc    bool
	Reorder bool
}

type OrderBy struct {
	Columns    []OrderByColumn
	Expression Expression
}

// Build build where clause
func (orderBy OrderBy) Build(builder Builder) {
	if orderBy.Expression != nil {
		orderBy.Expression.Build(builder)
	} else {
		for idx, column := range orderBy.Columns {
			if idx > 0 {
				builder.WriteByte(',')
			}

			builder.WriteQuoted(column.Column)
			if column.Desc {
				builder.WriteString(" DESC")
			}
		}
	}
}

// Name where clause name
func (limit Limit) Name() string {}

// MergeClause merge order by clause
func (limit Limit) MergeClause(clause *Clause) {}

// Name where clause name
func (limit Limit) Name() string {}

// MergeClause merge limit by clause
func (limit Limit) MergeClause(clause *Clause) {}
