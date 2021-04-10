package column

import "fmt"

type IntColumn struct {
	datatype string

	name   string
	length int
    size string

	allowNull bool
	isUnique  bool
	isPrimary   bool
	isAutoIncrement   bool
    isUnsigned bool

	defaultValue  string
	onUpdateValue string
	onDeleteValue string
}

func NewIntColumn(name string) *IntColumn {
    return &IntColumn{datatype: "int", name: name, allowNull: false, isUnique: false, length: 11, size: ""}
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
