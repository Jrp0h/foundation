package column

type EnumColumn struct {
	datatype string

	name   string
	values []string

	allowNull bool
	isUnique  bool
	isPrimary bool

	defaultValue  string
	onUpdateValue string
	onDeleteValue string

	alter bool
}

func NewEnumColumn(name string, values []string) *EnumColumn {

	return &EnumColumn{datatype: "ENUM", name: name, allowNull: false, isUnique: false, values: values, alter: false}
}

func (col *EnumColumn) Nullable() *EnumColumn {
	col.allowNull = true
	return col
}

func (col *EnumColumn) Unique() *EnumColumn {
	col.isUnique = true
	return col
}

func (col *EnumColumn) Primary() *EnumColumn {
	col.isPrimary = true
	return col
}

func (col *EnumColumn) Default(value string) *EnumColumn {

	isInvalid := true

	for _, v := range col.values {
		if v == value {
			isInvalid = false
			break
		}
	}

	if isInvalid {
		panic(value + " isn't valid option for " + col.name)
	}

	col.defaultValue = value
	return col
}

func (col *EnumColumn) OnUpdate(value string) *EnumColumn {
	col.onUpdateValue = value
	return col
}

func (col *EnumColumn) OnDelete(value string) *EnumColumn {
	col.onDeleteValue = value
	return col
}

func (col *EnumColumn) Alter() *EnumColumn {
	col.alter = true
	return col
}

func (col *EnumColumn) ToInsertSQL() string {
	sql := col.name + " " + col.datatype + "("

	for i, v := range col.values {
		sql += "'" + v + "'"
		if i != len(col.values)-1 {
			sql += ", "
		}
	}

	sql += ")"

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
		sql += " DEFAULT '" + col.defaultValue + "'"
	}

	return sql
}

func (col *EnumColumn) ToAlterSQL() string {
	if col.alter {
		return " MODIFY " + col.ToInsertSQL()
	}

	return " ADD " + col.ToInsertSQL()
}
