package column

import "fmt"

type ForeignIDColumn struct {
	datatype string

	name   string
	length int
	size   string

	allowNull  bool
	isUnique   bool
	isPrimary  bool
	isUnsigned bool

	defaultValue  string
	onUpdateValue string
	onDeleteValue string

	references string
	on         string

	alter bool
}

func NewForeignIDColumn(name string, references string, on string) *ForeignIDColumn {
	return &ForeignIDColumn{datatype: "INT", name: name, allowNull: false, isUnique: false, length: 11, references: references, on: on, alter: false}
}

func (col *ForeignIDColumn) Tiny() *ForeignIDColumn {
	return col.Size("TINY")
}

func (col *ForeignIDColumn) Small() *ForeignIDColumn {
	return col.Size("SMALL")
}

func (col *ForeignIDColumn) Medium() *ForeignIDColumn {
	return col.Size("MEDIUM")
}

func (col *ForeignIDColumn) Big() *ForeignIDColumn {
	return col.Size("BIG")
}

func (col *ForeignIDColumn) Size(size string) *ForeignIDColumn {
	switch size {
	case "TINY", "SMALL", "", "MEDIUM", "BIG":
		col.size = size
		break
	default:
		panic(size + " is not a valid integer size!, allowed: TINY, SMALL, MEDIUM, BIG and ''(Empty string)")
	}

	return col
}

func (col *ForeignIDColumn) Length(length int) *ForeignIDColumn {
	col.length = length
	return col
}

func (col *ForeignIDColumn) Nullable() *ForeignIDColumn {
	col.allowNull = true
	return col
}

func (col *ForeignIDColumn) Unique() *ForeignIDColumn {
	col.isUnique = true
	return col
}

func (col *ForeignIDColumn) Primary() *ForeignIDColumn {
	col.isPrimary = true
	return col
}

func (col *ForeignIDColumn) Unsigned() *ForeignIDColumn {
	col.isUnsigned = true
	return col
}

func (col *ForeignIDColumn) Default(value int) *ForeignIDColumn {
	col.defaultValue = fmt.Sprint(value)
	return col
}

func (col *ForeignIDColumn) OnUpdate(value string) *ForeignIDColumn {
	col.onUpdateValue = value
	return col
}

func (col *ForeignIDColumn) OnDelete(value string) *ForeignIDColumn {
	col.onDeleteValue = value
	return col
}

func (col *ForeignIDColumn) Alter() *ForeignIDColumn {
	col.alter = true
	return col
}

func (col *ForeignIDColumn) ToInsertSQL() string {
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

	if col.defaultValue != "" {
		sql += " DEFAULT " + col.defaultValue
	}

	sql += ",\n\tFOREIGN KEY (" + col.name + ") REFERENCES " + col.references + "(" + col.on + ")"

	if col.onUpdateValue != "" {
		sql += " ON UPDATE " + col.onUpdateValue
	}

	if col.onDeleteValue != "" {
		sql += " ON DELETE " + col.onDeleteValue
	}

	return sql
}

func (col *ForeignIDColumn) ToAlterSQL() string {
	if col.alter {
		return " MODIFY " + col.ToInsertSQL()
	}

	return " ADD " + col.ToInsertSQL()
}
