package main

import (
    "database/sql"
	"fmt"
	. "foundation/table"
	_ "github.com/go-sql-driver/mysql"
)

type Migration interface
{
    Up()
    Down()
}

func main() {
    companies := CreateTable("companies", func (table *Table) {
        table.ID()
        table.String("name")
        table.String("address")
    })

    users := CreateTable("users", func (table *Table) {
        table.ID()
        table.String("email")
        table.Int("age")
        table.Enum("roles", []string{"Owner", "Maintainer", "Developer", "Guest"})
        table.ForeignID("company_id", "companies", "id")
        table.Bool("is_male").Default(true)
        table.Timestamps()
    })

    db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/go_db_test")

    if err != nil {
        fmt.Println("Shit fucked up")
        fmt.Println(err.Error())
        return 
    }

    defer db.Close()

    stmt, stmtError := db.Prepare(companies)

    if stmtError != nil {
        fmt.Println("Companies statement")
        fmt.Println(stmtError.Error())
        return
    }

    fmt.Println("Running: ")
    fmt.Println(companies)
    fmt.Println()
    stmt.Exec()

    stmt, stmtError = db.Prepare(users)

    if stmtError != nil {
        fmt.Println("Users statement")
        fmt.Println(stmtError.Error())
        return
    }

    fmt.Println("Running: ")
    fmt.Println(users)
    stmt.Exec()
}
