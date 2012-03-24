package main

type Function func(*Stack) VType
type Table struct {
	functions map[string] Function
	aliases   map[string] string
}

func (t *Table) get(name string) Function {
	if t.has(name) {
		f, fp := t.functions[name]
		if fp {
			return f
		} else {
			a, _ := t.aliases[name]
			return t.functions[a]
		}
	}
	return nil
}

func (t *Table) has(name string) bool {
	_, f := t.functions[name]
	_, a := t.aliases[name]
	return f || a
}

func (t *Table) define(name string, f Function) {
	t.functions[name] = f
}

func (t *Table) alias(from, to string) {
	t.aliases[from] = to
}


func NewTable() *Table {
	tbl := new(Table)

	f := map[string] Function { }
	tbl.functions = f

	a := map[string] string { }
	tbl.aliases = a

	return tbl
}
