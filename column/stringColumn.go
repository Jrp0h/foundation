package column

import "fmt"

type StringColumn struct {
	datatype string

	name   string
	length int

	allowNull bool
	isUnique  bool
	isPrimary   bool

	defaultValue  string
	onUpdateValue string
	onDeleteValue string
}

func NewStringColumn(name string) *StringColumn {
	return &StringColumn{datatype: "VARCHAR", name: name, allowNull: false, isUnique: false, length: 128}
}

func (col *StringColumn) Length(length int) *StringColumn {
	col.length = length
	return col
}

func (col *StringColumn) Nullable() *StringColumn {
	col.allowNull = true
	return col
}

func (col *StringColumn) Unique() *StringColumn {
	col.isUnique = true
	return col
}

func (col *StringColumn) Primary() *StringColumn {
	col.isPrimary = true
	return col
}

func (col *StringColumn) Default(value string) *StringColumn {
	col.defaultValue = value
	return col
}

func (col *StringColumn) OnUpdate(value string) *StringColumn {
	col.onUpdateValue = value
	return col
}

func (col *StringColumn) OnDelete(value string) *StringColumn {
	col.onDeleteValue = value
	return col
}

func (col *StringColumn) ToSQL() string {
    sql := col.name + " " + col.datatype + "(" + fmt.Sprint(col.length) + ")"

    if !col.allowNull {
        sql += " NOT NULL"
    }

    if col.isUnique {
        sql += " UNIQUE"
    }

    if col.isPrimary {
        sql += " PRIMARY KEY"
    }

    if col.onUpdateValue != "" {
        sql += " ON UPDATE " + col.onUpdateValue
    }

    if col.onDeleteValue != "" {
        sql += " ON DELETE " + col.onDeleteValue
    }

    if col.defaultValue != "" {
        sql += " DEFAULT '" + col.defaultValue + "'"
    }

    return sql
}
