package main

import (
	. "github.com/Jrp0h/foundation/schema"
	. "github.com/Jrp0h/foundation/table"
)

func main() {
	RunMigrations(Credentials{
		Username: "root",
		Password: "root",
		Ip:       "127.0.0.1",
		Port:     3306,
		Name:     "foundation_test",
	}, up, down)
}

func up(schema *Schema) {
	schema.CreateTableIfNotExists("companies", func(table *Table) {
		table.ID()
		table.String("name")
		table.String("address")
	})

	schema.CreateTableIfNotExists("users", func(table *Table) {
		table.ID()
		table.String("email")
		table.Int("age").Nullable()
		table.Enum("roles", []string{"Owner", "Maintainer", "Developer", "Guest"}).Default("Guest")
		table.ForeignID("company_id", "companies", "id")
		table.Bool("is_male").Default(true)
		table.Timestamps()
	})

	schema.AlterTable("users", func(table *Table) {
		table.Enum("is_male", []string{"Male", "Female"}).Default("Male").Alter()
		table.Enum("Gender", []string{"Male", "Female"}).Default("Female")
		table.DropColumn("age")
	})

}

func down(schema *Schema) {
	schema.DropTableIfExists("users")
	schema.DropTableIfExists("companies")
}
