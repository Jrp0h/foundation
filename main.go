package main

import (
    // "fmt"
    // . "foundation/column"
    . "foundation/table"
)

type Migration interface
{
    Up()
    Down()
}

func main() {
    CreateTable("users", func (table *Table) {
        table.ID()
        table.String("email")
        table.Int("age")
        table.Enum("roles", []string{"Owner", "Maintainer", "Developer", "Guest"})
        table.ForeignID("company_id", "companies", "id")
        table.Bool("is_male").Default(true)
        table.Timestamps()
    })
}
