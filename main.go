package main

import (
    . "foundation/table"
    . "foundation/schema"
)

func main() {
    RunMigrations(Credentials{
        Username: "root",
        Password: "",
        Ip: "127.0.0.1",
        Port: 3306,
        Name: "go_db_test",
    }, up, down);
}

func up(schema *Schema) {
    schema.CreateTable("companies", func (table *Table) {
        table.ID()
        table.String("name")
        table.String("address")
    })

    schema.CreateTable("users", func (table *Table) {
        table.ID()
        table.String("email")
        table.Int("age").Nullable()
        table.Enum("roles", []string{"Owner", "Maintainer", "Developer", "Guest"})
        table.ForeignID("company_id", "companies", "id")
        table.Bool("is_male").Default(true)
        table.Timestamps()
    })

}

func down(schema *Schema) {
    schema.DropIfExists("companies")
    schema.DropIfExists("users")
}
