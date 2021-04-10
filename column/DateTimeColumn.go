package column

type DateTimeColumn struct {
	datatype string

	name   string

	allowNull bool
	isUnique  bool
	isPrimary   bool
	isAutoIncrement   bool

	defaultValue  string
	onUpdateValue string
	onDeleteValue string
}

func NewDateTimeColumn(name string) *DateTimeColumn {
	return &DateTimeColumn{datatype: "datetime", name: name, allowNull: false, isUnique: false}
}

func (col *DateTimeColumn) Nullable() *DateTimeColumn {
	col.allowNull = true
	return col
}

func (col *DateTimeColumn) Unique() *DateTimeColumn {
	col.isUnique = true
	return col
}

func (col *DateTimeColumn) Primary() *DateTimeColumn {
	col.isPrimary = true
	return col
}

func (col *DateTimeColumn) Default(value string) *DateTimeColumn {
	col.defaultValue = value
	return col
}

func (col *DateTimeColumn) OnUpdate(value string) *DateTimeColumn {
	col.onUpdateValue = value
	return col
}

func (col *DateTimeColumn) OnDelete(value string) *DateTimeColumn {
	col.onDeleteValue = value
	return col
}

func (col *DateTimeColumn) ToSQL() string {
    sql := col.name + " " + col.datatype

    if !col.allowNull {
        sql += " NOT NULL"
    }

    if col.isUnique {
        sql += " UNIQUE"
    }

    if col.isPrimary {
        sql += " PRIMARY KEY"
    }

    if col.isAutoIncrement {
        sql += " AUTO_INCREMENT"
    }

    if col.onUpdateValue != "" {
        sql += " ON UPDATE " + col.onUpdateValue
    }

    if col.onDeleteValue != "" {
        sql += " ON DELETE " + col.onDeleteValue
    }

    if col.defaultValue != "" {
        sql += " DEFAULT " + col.defaultValue
    }

    return sql
}
