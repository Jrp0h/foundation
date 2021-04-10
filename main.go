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
    });
}
