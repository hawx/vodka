package main

type Function func(*Stack, *Table) VType
type Table struct {
	native    map[string]Tokens
	functions map[string]Function
	aliases   map[string]string
}

func (t *Table) String() string {
	s := "Native:\n"
	for i, ts := range t.native {
		s += "  " + i + " " + ts.String() + "\n"
	}

	s += "\nFunctions:\n"
	for i, _ := range t.functions {
		s += "  " + i
	}

	return s + "\n"
}

func (t *Table) Defined() string {
	s := ""
	for n, _ := range t.native {
		s += n + " "
	}
	for n, _ := range t.functions {
		s += n + " "
	}
	for n, _ := range t.aliases {
		s += n + " "
	}
	return s
}

func (t *Table) call(name string, s *Stack) VType {
	var e VType = VNil()
	if n, np := t.native[name]; np {
		_, _, e = Run(&n, s, t)
	} else if f, fp := t.functions[name]; fp {
		e = f(s, t)
	} else {
		a, _ := t.aliases[name]
		return t.call(a, s)
	}
	return e
}

func (t *Table) has(name string) bool {
	_, n := t.native[name]
	_, f := t.functions[name]
	_, a := t.aliases[name]
	return n || f || a
}

func (t *Table) define(name string, f Function) {
	t.functions[name] = f
}

func (t *Table) defineNative(name string, ts *Tokens) {
	t.native[name] = *ts
}

func (t *Table) alias(from, to string) {
	t.aliases[from] = to
}

func NewTable() *Table {
	tbl := new(Table)

	n := map[string]Tokens{}
	tbl.native = n

	f := map[string]Function{}
	tbl.functions = f

	a := map[string]string{}
	tbl.aliases = a

	return tbl
}
