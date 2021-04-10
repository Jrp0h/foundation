package main

import (
    // "fmt"
    // . "foundation/column"
    . "foundation/table"
)

func main() {
    CreateTable("users", func (table *Table) {
        table.ID()
        table.String("email")
        table.Int("age")
        table.Enum("roles", []string{"Owner", "Maintainer", "Developer", "Guest"})
        table.ForeginID("company_id", "companies", "id")
        table.Timestamps()
    })
}
