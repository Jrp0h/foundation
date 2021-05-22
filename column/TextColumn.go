package column

type TextColumn struct {
	datatype string
	size     string

	name string

	allowNull bool
	isUnique  bool
	isPrimary bool

	defaultValue  string
	onUpdateValue string
	onDeleteValue string

	alter bool
}

func NewTextColumn(name string) *TextColumn {
	return &TextColumn{datatype: "TEXT", name: name, allowNull: false, isUnique: false, size: "", alter: false}
}

func (col *TextColumn) Tiny() *TextColumn {
	return col.Size("TINY")
}

func (col *TextColumn) Medium() *TextColumn {
	return col.Size("MEDIUM")
}

func (col *TextColumn) Long() *TextColumn {
	return col.Size("LONG")
}

func (col *TextColumn) Size(size string) *TextColumn {
	switch size {
	case "TINY", "", "MEDIUM", "LONG":
		col.size = size
		break
	default:
		panic(size + " is not a valid text type!, allowed: TINY, MEDIUM, LONG and ''(Empty string)")
	}

	return col
}

func (col *TextColumn) Nullable() *TextColumn {
	col.allowNull = true
	return col
}

func (col *TextColumn) Unique() *TextColumn {
	col.isUnique = true
	return col
}

func (col *TextColumn) Primary() *TextColumn {
	col.isPrimary = true
	return col
}

func (col *TextColumn) Default(value string) *TextColumn {
	col.defaultValue = value
	return col
}

func (col *TextColumn) OnUpdate(value string) *TextColumn {
	col.onUpdateValue = value
	return col
}

func (col *TextColumn) OnDelete(value string) *TextColumn {
	col.onDeleteValue = value
	return col
}

func (col *TextColumn) ToInsertSQL() string {
	sql := col.name + " " + col.size + col.datatype

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

func (col *TextColumn) Alter() *TextColumn {
	col.alter = true
	return col
}

func (col *TextColumn) ToAlterSQL() string {
	if col.alter {
		return " MODIFY " + col.ToInsertSQL()
	}

	return " ADD IF NOT EXISTS " + col.ToInsertSQL()
}
