package column;

type Column struct {
    name string
    allowNull bool
    isUnique bool
}

func (column *Column) setName(name string) (*Column) {
    column.name = name;
    return column;
}

func (column *Column) nullable(nullable bool) (*Column) {
    column.allowNull = nullable;
    return column;
}

func (column *Column) unique(unique bool) (*Column) {
    column.isUnique = unique;
    return column;
}

func NewColumn(name string) (*Column) {

    return &Column{name: name, allowNull: false, isUnique: false}
}

