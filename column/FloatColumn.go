package column

import "fmt"

type FloatColumn struct {
	datatype string

	name   string
    size int
    d int

	allowNull bool
	isUnique  bool
	isPrimary   bool
	isAutoIncrement   bool
    isUnsigned bool

	defaultValue  string
	onUpdateValue string
	onDeleteValue string
}

func NewFloatColumn(name string, size int, d int) *FloatColumn {
    return &FloatColumn{datatype: "FLOAT", name: name, allowNull: false, isUnique: false, size: size, d: d}
}

func (col *FloatColumn) Nullable() *FloatColumn {
	col.allowNull = true
	return col
}

func (col *FloatColumn) Unique() *FloatColumn {
	col.isUnique = true
	return col
}

func (col *FloatColumn) Primary() *FloatColumn {
	col.isPrimary = true
	return col
}

func (col *FloatColumn) AutoIncrement() *FloatColumn {
	col.isAutoIncrement = true
	return col
}

func (col *FloatColumn) Unsigned() *FloatColumn {
	col.isUnsigned = true
	return col
}

func (col *FloatColumn) Default(value int) *FloatColumn {
	col.defaultValue = fmt.Sprint(value)
	return col
}

func (col *FloatColumn) OnUpdate(value string) *FloatColumn {
	col.onUpdateValue = value
	return col
}

func (col *FloatColumn) OnDelete(value string) *FloatColumn {
	col.onDeleteValue = value
	return col
}

func (col *FloatColumn) ToInsertSQL() string {
    sql := col.name + " " + col.datatype + "(" + fmt.Sprint(col.size) + ", " + fmt.Sprint(col.d) + ")"

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
