package table

import (
    "fmt"
	. "foundation/column"
)

type SQLable interface {
    ToSQL() string
}

type Table struct {
    name string
    columns []*SQLable
}

func CreateTable(name string, closure func(*Table)) {
    table := Table {name: name}

    closure(&table);

    sql := "CREATE TABLE " + table.name + " (\n"

    for i, col := range table.columns {
        sql += SQLable(*col).ToSQL()

        if i != len(table.columns) - 1 {
            sql += ", \n"
        }

    }

    sql += ");"

    fmt.Println(sql)
}

func (table *Table) String(name string) (*StringColumn) {
    s := NewStringColumn(name)
    sqlable := SQLable(s)
    table.columns = append(table.columns, &sqlable)
    return s
}
