# Foundation

Foundation is a Laravel Eloquent inspired sql migration handle(?) made for fun.

## Documentation

To create a table:

```golang
CreateTable("name", func(table *Table) {
    // Add columns here
    table.ID()
    table.String("email")
    table.Int("age")
    table.Enum("roles", []string{"Owner", "Maintainer", "Developer", "Guest"})
    table.ForeignID("company_id", "companies", "id")
    table.Bool("is_male").Default(true)
    table.Timestamps()
})
```

### Columns

* String
* DateTime
* Int
* Text
* ForeignID
* Enum
* Bool

#### String

Takes in a column name

Translates to a varchar

Methods:

Length - Sets the length of the string
Nullable - Removes NOT NULL from final string
Unique - Adds UNIQUE to final string
Primary - Adds PRIMARY KEY to final string
Default - Adds a default value when no value is inserted in the database
OnUpdate - Sql on update
OnDelete - Sql on delete

#### DateTime

Takes in a column name

Translates to a datetime

Methods:

Nullable - Removes NOT NULL from final string
Unique - Adds UNIQUE to final string
Primary - Adds PRIMARY KEY to final string
Default - Adds a default value when no value is inserted in the database
OnUpdate - Sql on update
OnDelete - Sql on delete

#### Int

Takes in a column name

Translates to a Int or TINYINT, SMALLINT, MEDIUMINT, BIGINT depending

Methods:

Length - Sets the length of the int
Tiny - Makes it a TINYINT
Small - Makes it a SMALLINT
Medium - Makes it a MEDIUMINT
Big - Makes it a BIGINT
Unsigned - Makes it unsigned
Nullable - Removes NOT NULL from final string
Unique - Adds UNIQUE to final string
Primary - Adds PRIMARY KEY to final string
AutoIncrement - Adds AUTO_INCREMENT to final string
Default - Adds a default value when no value is inserted in the database
OnUpdate - Sql on update
OnDelete - Sql on delete

#### Text

Takes in a column name

Translates to a text or SMALLTEXT, MEDIUMTEXT, LONGTEXT depending

Methods:

Length - Sets the length of the string
Small - Makes it a SMALLTEXT
Medium - Makes it a MEDIUMTEXT
Long - Makes it a LONGTEXT
Nullable - Removes NOT NULL from final string
Unique - Adds UNIQUE to final string
Primary - Adds PRIMARY KEY to final string
Default - Adds a default value when no value is inserted in the database
OnUpdate - Sql on update
OnDelete - Sql on delete

#### ForeignID

Takes in a column name, what it references(table) and on what (field)

Translates to a Int or TINYINT, SMALLINT, MEDIUMINT, BIGINT depending

Automatically set to UNSIGNED BIGINT

Methods:

Length - Sets the length of the int
Tiny - Makes it a TINYINT
Small - Makes it a SMALLINT
Medium - Makes it a MEDIUMINT
Big - Makes it a BIGINT
Nullable - Removes NOT NULL from final string
Unsigned - Makes it unsigned
Unique - Adds UNIQUE to final string
Primary - Adds PRIMARY KEY to final string
Default - Adds a default value when no value is inserted in the database
OnUpdate - Sql on update
OnDelete - Sql on delete

#### Enum

Takes in a column name and a string array of values

Translates to a ENUM

Methods:

Nullable - Removes NOT NULL from final string
Unique - Adds UNIQUE to final string
Primary - Adds PRIMARY KEY to final string
Default - Adds a default value when no value is inserted in the database
OnUpdate - Sql on update
OnDelete - Sql on delete

#### Bool

Takes in a column name and a string array of values

Translates to a Bool

Methods:

Nullable - Removes NOT NULL from final string
Unique - Adds UNIQUE to final string
Primary - Adds PRIMARY KEY to final string
Default - Adds a default value when no value is inserted in the database
OnUpdate - Sql on update
OnDelete - Sql on delete

### Helper methods

```golang
table.ID()

// Same as
table.Int("id").AutoIncrement().Primary().Unsigned().Big()
```

```golang
table.Timestamps()

// Same as
table.DateTime("created_at").Default("NOW()")
table.DateTime("updated_at").Default("NOW()").OnUpdate("NOW()")
```

```golang
table.TinyText(name)

// Same as
table.Text(name).Tiny()
```


```golang
table.MediumText(name)

// Same as
table.Text(name).Medium()
```

```golang
table.LongText(name)

// Same as
table.Text(name).Long()
```

```golang
table.TinyInt(name)

// Same as
table.Int(name).Tiny()
```

```golang
table.SmallInt(name)

// Same as
table.Int(name).Small()
```

```golang
table.MediumInt(name)

// Same as
table.Int(name).Medium()
```

```golang
table.BigInt(name)

// Same as
table.Int(name).Big()
```
