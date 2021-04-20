package schema

import (
	"database/sql"
	"fmt"
    "os"
    _ "github.com/go-sql-driver/mysql"
    . "foundation/table"
)

type Credentials struct {
    Username string
    Password string
    Ip string
    Port int
    Name string
}

type Schema struct {
    credentials Credentials
}

func (c Credentials) getConnectionString() string {
    return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", c.Username, c.Password, c.Ip, c.Port, c.Name)
}

func (s* Schema) sqlRun(query string) {
    db, err := sql.Open("mysql", s.credentials.getConnectionString())
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

func RunMigrations(credentials Credentials, up func(*Schema), down func(*Schema)) {

    s := Schema{credentials: credentials}

    args := os.Args[1:]

    if len(args) == 1 {
        if args[0] == "up" {
            up(&s)
        } else if args[0] == "down" {
            down(&s)
        } else {
            fmt.Println("Invalid argument")
        }
    } else {
        fmt.Println("One argument required")
    }
}

func (s *Schema) CreateTable(name string, closure func(*Table)) string {
    table := Table {Name: name}

    closure(&table);

    sql := "CREATE TABLE " + table.Name + " (\n"

    for i, col := range table.Columns {
        sql += "\t" + SQLable(*col).ToInsertSQL()

        if i != len(table.Columns) - 1 {
            sql += ","
        }

        sql += "\n"

    }

    sql += ");"
    
    fmt.Println("Creating table " + name)
    s.sqlRun(sql)
    fmt.Println(name + " created")

    return sql
}

func (s* Schema) DropIfExists(name string) string {
    // Drop table
    sql := "DROP TABLE IF EXISTS " + name + ";"

    fmt.Println("Dropping table " + name)
    s.sqlRun(sql)
    fmt.Println(name + " table dropped")

    return sql
}
