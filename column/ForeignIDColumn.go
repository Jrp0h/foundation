package column

import "fmt"

type ForeginIDColumn struct {
	datatype string

	name string
	length int
    size string

	allowNull bool
	isUnique  bool
	isPrimary bool
	isAutoIncrement bool
    isUnsigned bool

	defaultValue  string
	onUpdateValue string
	onDeleteValue string

    references string
    on string
}

func NewForeginIDColumn(name string, references string, on string) *ForeginIDColumn {
    return &ForeginIDColumn{datatype: "int", name: name, allowNull: false, isUnique: false, length: 11, references: references, on: on}
}

func (col *ForeginIDColumn) Size(size string) *ForeginIDColumn {
    switch size {
        case "TINY", "SMALL", "", "MEDIUM", "BIG":
            col.size = size
            break
        default:
            panic(size + " is not a valid integer size!, allowed: TINY, SMALL, MEDIUM, BIG and ''(Empty string)")
    }

	return col
}

func (col *ForeginIDColumn) Length(length int) *ForeginIDColumn {
	col.length = length
	return col
}

func (col *ForeginIDColumn) Nullable() *ForeginIDColumn {
	col.allowNull = true
	return col
}

func (col *ForeginIDColumn) Unique() *ForeginIDColumn {
	col.isUnique = true
	return col
}

func (col *ForeginIDColumn) Primary() *ForeginIDColumn {
	col.isPrimary = true
	return col
}

func (col *ForeginIDColumn) AutoIncrement() *ForeginIDColumn {
	col.isAutoIncrement = true
	return col
}

func (col *ForeginIDColumn) Unsigned() *ForeginIDColumn {
	col.isUnsigned = true
	return col
}

func (col *ForeginIDColumn) Default(value int) *ForeginIDColumn {
	col.defaultValue = fmt.Sprint(value)
	return col
}

func (col *ForeginIDColumn) OnUpdate(value string) *ForeginIDColumn {
	col.onUpdateValue = value
	return col
}

func (col *ForeginIDColumn) OnDelete(value string) *ForeginIDColumn {
	col.onDeleteValue = value
	return col
}

func (col *ForeginIDColumn) ToInsertSQL() string {
    sql := col.name + " " + col.datatype + "(" + fmt.Sprint(col.length) + ")"

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

    sql += ",\nFOREGIN KEY (" + col.name + ") REFERENCES " + col.references + "(" + col.on + ")"

    return sql
}
