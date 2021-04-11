package table

import (
    "database/sql"
	. "foundation/column"
    _ "github.com/go-sql-driver/mysql"
    "fmt"
)

type SQLable interface {
    ToInsertSQL() string
    // ToAlterSQL(tableName string)
}

type Table struct {
    name string
    columns []*SQLable
}

func sqlRun(query string) {
    db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/go_db_test")
    if err != nil {
        fmt.Println("Couldn't open database")
        panic("Couldn't open database")
    }

    defer db.Close()

    stmt, stmtError := db.Prepare(query)

    if stmtError != nil {
        fmt.Println("Error prepering statment:")
        fmt.Println(query)
        fmt.Println(stmtError.Error())
        panic("error prepering statment")
    }

    _, execError := stmt.Exec()

    if execError != nil {
        fmt.Println("Error running statment:")
        fmt.Println(query)
        fmt.Println(execError.Error())
        panic("Error running statment")
    }
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

    sql += ");"

    sqlRun(sql)

    return sql
}

func DropIfExists(name string) string {
    // Drop table
    sql := "DROP TABLE IF EXISTS " + name + ";"
    sqlRun(sql)

    return sql
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

