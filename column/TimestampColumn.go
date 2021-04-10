package column

type TimestampColumn struct {
	datatype string

	name   string

	allowNull bool
	isUnique  bool
	isPrimary   bool

	defaultValue  string
	onUpdateValue string
	onDeleteValue string
}

func NewTimestampColumn(name string) *TimestampColumn {
	return &TimestampColumn{datatype: "TIMESTAMP", name: name, allowNull: false, isUnique: false}
}

func (col *TimestampColumn) Nullable() *TimestampColumn {
	col.allowNull = true
	return col
}

func (col *TimestampColumn) Unique() *TimestampColumn {
	col.isUnique = true
	return col
}

func (col *TimestampColumn) Primary() *TimestampColumn {
	col.isPrimary = true
	return col
}

func (col *TimestampColumn) Default(value string) *TimestampColumn {
	col.defaultValue = value
	return col
}

func (col *TimestampColumn) OnUpdate(value string) *TimestampColumn {
	col.onUpdateValue = value
	return col
}

func (col *TimestampColumn) OnDelete(value string) *TimestampColumn {
	col.onDeleteValue = value
	return col
}

func (col *TimestampColumn) ToInsertSQL() string {
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
