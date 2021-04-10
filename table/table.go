package table

import (
    // "fmt"
	. "foundation/column"
)

type SQLable interface {
    ToInsertSQL() string
    // ToAlterSQL(tableName string)
}

type Table struct {
    name string
    columns []*SQLable
}

func CreateTable(name string, closure func(*Table)) string {
    table := Table {name: name}

    closure(&table);

    sql := "CREATE TABLE " + table.name + " (\n"

    for i, col := range table.columns {
        sql += "\t" + SQLable(*col).ToInsertSQL()

        if i != len(table.columns) - 1 {
            sql += ","
        }

        sql += "\n"

    }

    return sql + ");"
}

func DropIfExists(name string) string {
    // Drop table
    return "DROP IF EXISTS " + name + ";"
}

// Default datatypes
func (table *Table) String(name string) (*StringColumn) {
    s := NewStringColumn(name)
    sqlable := SQLable(s)
    table.columns = append(table.columns, &sqlable)
    return s
}

func (table *Table) Int(name string) (*IntColumn) {
    s := NewIntColumn(name)
    sqlable := SQLable(s)
    table.columns = append(table.columns, &sqlable)
    return s
}

func (table *Table) Timestamp(name string) (*TimestampColumn) {
    s := NewTimestampColumn(name)
    sqlable := SQLable(s)
    table.columns = append(table.columns, &sqlable)
    return s
}

func (table *Table) Bool(name string) (*BoolColumn) {
    s := NewBoolColumn(name)
    sqlable := SQLable(s)
    table.columns = append(table.columns, &sqlable)
    return s
}

func (table *Table) Enum(name string, values []string) (*EnumColumn) {
    s := NewEnumColumn(name, values)
    sqlable := SQLable(s)
    table.columns = append(table.columns, &sqlable)
    return s
}

func (table *Table) Text(name string) (*TextColumn) {
    s := NewTextColumn(name)
    sqlable := SQLable(s)
    table.columns = append(table.columns, &sqlable)
    return s
}

func (table *Table) ForeignID(name string, references string, on string) (*ForeignIDColumn) {
    s := NewForeignIDColumn(name, references, on)
    s.Unsigned().Size("BIG")
    sqlable := SQLable(s)
    table.columns = append(table.columns, &sqlable)
    return s
}


// Helper datatypes
func (table *Table) ID() (*IntColumn) {
    s := NewIntColumn("id")
    s.AutoIncrement().Primary().Unique().Unsigned().Big()
    sqlable := SQLable(s)
    table.columns = append(table.columns, &sqlable)
    return s
}

func (table *Table) Timestamps() {
    table.Timestamp("created_at").Default("NOW()")
    table.Timestamp("updated_at").Default("NOW()").OnUpdate("NOW()")
}

func (table *Table) TinyText(name string) (*TextColumn) {
    return table.Text(name).Tiny()
}

func (table *Table) MediumText(name string) (*TextColumn) {
    return table.Text(name).Medium()
}

func (table *Table) LongText(name string) (*TextColumn) {
    return table.Text(name).Long()
}

func (table *Table) TinyInt(name string) (*IntColumn) {
    return table.Int(name).Tiny()
}

func (table *Table) SmallInt(name string) (*IntColumn) {
    return table.Int(name).Small()
}

func (table *Table) MediumInt(name string) (*IntColumn) {
    return table.Int(name).Medium()
}

func (table *Table) BigInt(name string) (*IntColumn) {
    return table.Int(name).Big()
}
