package main

import (
    "os"
    "fmt"
    . "foundation/table"
)

func main() {
    args := os.Args[1:]

    if len(args) == 1 {
        if args[0] == "up" {
            up()
        } else if args[0] == "down" {
            down()
        } else {
            fmt.Println("Invalid argument")
        }
    } else {
        fmt.Println("One argument required")
    }
}

func up() {
    CreateTable("companies", func (table *Table) {
        table.ID()
        table.String("name")
        table.String("address")
    })

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

func down() {
    DropIfExists("users")
    DropIfExists("companies")
}
