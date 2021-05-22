package column

type DropColumn struct {
	name string
}

func NewDropColumn(name string) *DropColumn {
	return &DropColumn{name: name}
}

func (col *DropColumn) ToInsertSQL() string {
	panic("You should not drop a column in a insert query")
}

func (col *DropColumn) ToAlterSQL() string {
	return " DROP COLUMN " + col.name
}
