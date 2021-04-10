package main

import (
    // "fmt"
    // . "foundation/column"
    . "foundation/table"
)

func main() {
    CreateTable("users", func (table *Table) {
       table.String("email").Primary()
       table.String("username").Unique()
       table.String("password").Default("$2y$12$lmODarB31RKvIKThIBpTbOIH8ZcrPkRiaKdVQ020z/IxhyX3o5Snu")
       table.String("firstname").Nullable().Length(255)
       table.String("foreginKey").OnDelete("cascade").OnUpdate("cascade")
    });
}
