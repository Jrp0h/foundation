package column

import "fmt"

type IntColumn struct {
	datatype string

	name   string
	length int
	size   string

	allowNull       bool
	isUnique        bool
	isPrimary       bool
	isAutoIncrement bool
	isUnsigned      bool

	defaultValue  string
	onUpdateValue string
	onDeleteValue string

	alter bool
}

func NewIntColumn(name string) *IntColumn {
	return &IntColumn{datatype: "INT", name: name, allowNull: false, isUnique: false, length: 11, size: "", alter: false}
}

func (col *IntColumn) Tiny() *IntColumn {
	return col.Size("TINY")
}

func (col *IntColumn) Small() *IntColumn {
	return col.Size("SMALL")
}

func (col *IntColumn) Medium() *IntColumn {
	return col.Size("MEDIUM")
}

func (col *IntColumn) Big() *IntColumn {
	return col.Size("BIG")
}

func (col *IntColumn) Size(size string) *IntColumn {
	switch size {
	case "TINY", "SMALL", "", "MEDIUM", "BIG":
		col.size = size
		break
	default:
		panic(size + " is not a valid integer size!, allowed: TINY, SMALL, MEDIUM, BIG and ''(Empty string)")
	}

	return col
}

func (col *IntColumn) Length(length int) *IntColumn {
	col.length = length
	return col
}

func (col *IntColumn) Nullable() *IntColumn {
	col.allowNull = true
	return col
}

func (col *IntColumn) Unique() *IntColumn {
	col.isUnique = true
	return col
}

func (col *IntColumn) Primary() *IntColumn {
	col.isPrimary = true
	return col
}

func (col *IntColumn) AutoIncrement() *IntColumn {
	col.isAutoIncrement = true
	return col
}

func (col *IntColumn) Unsigned() *IntColumn {
	col.isUnsigned = true
	return col
}

func (col *IntColumn) Default(value int) *IntColumn {
	col.defaultValue = fmt.Sprint(value)
	return col
}

func (col *IntColumn) OnUpdate(value string) *IntColumn {
	col.onUpdateValue = value
	return col
}

func (col *IntColumn) OnDelete(value string) *IntColumn {
	col.onDeleteValue = value
	return col
}

func (col *IntColumn) Alter() *IntColumn {
	col.alter = true
	return col
}

func (col *IntColumn) ToInsertSQL() string {
	sql := col.name + " " + col.size + col.datatype + "(" + fmt.Sprint(col.length) + ")"

	if col.isUnsigned {
		sql += " UNSIGNED"
	}

	if !col.allowNull {
		sql += " NOT NULL"
	}

	if col.isUnique {
		sql += " UNIQUE"
	}

	if col.isPrimary {
		sql += " PRIMARY KEY"
	}

	if col.isAutoIncrement {
		sql += " AUTO_INCREMENT"
	}

	if col.onUpdateValue != "" {
		sql += " ON UPDATE " + col.onUpdateValue
	}

	if col.onDeleteValue != "" {
		sql += " ON DELETE " + col.onDeleteValue
	}

	if col.defaultValue != "" {
		sql += " DEFAULT " + col.defaultValue
	}

	return sql
}

func (col *IntColumn) ToAlterSQL() string {
	if col.alter {
		return " MODIFY " + col.ToInsertSQL()
	}

	return " ADD IF NOT EXISTS " + col.ToInsertSQL()
}
