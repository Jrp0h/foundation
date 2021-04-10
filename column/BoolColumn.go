package column

import "fmt"

type BoolColumn struct {
	datatype string

	name   string

	allowNull bool
	isUnique  bool
	isPrimary   bool

	defaultValue  string
	onUpdateValue string
	onDeleteValue string
}

func NewBoolColumn(name string) *BoolColumn {
    return &BoolColumn{datatype: "BOOL", name: name, allowNull: false, isUnique: false}
}

func (col *BoolColumn) Nullable() *BoolColumn {
	col.allowNull = true
	return col
}

func (col *BoolColumn) Unique() *BoolColumn {
	col.isUnique = true
	return col
}

func (col *BoolColumn) Primary() *BoolColumn {
	col.isPrimary = true
	return col
}

// I'm totaly not new to golang
func (col *BoolColumn) Default(value bool) *BoolColumn {
    if value {
        col.defaultValue = fmt.Sprint(1)
    } else  {
        col.defaultValue = fmt.Sprint(0)
    }
	return col
}

func (col *BoolColumn) OnUpdate(value string) *BoolColumn {
	col.onUpdateValue = value
	return col
}

func (col *BoolColumn) OnDelete(value string) *BoolColumn {
	col.onDeleteValue = value
	return col
}

func (col *BoolColumn) ToInsertSQL() string {
    sql := col.name + " " + col.datatype

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
        sql += " DEFAULT " + col.defaultValue
    }

    return sql
}
