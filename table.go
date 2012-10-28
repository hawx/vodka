package main

import (
	p "./parser"
	"./types"
	s "./stack"

	"./types/vnil"
)

type Function func(*s.Stack, *Table) types.VType
type Table struct {
	native    map[string]p.Tokens
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

func (t *Table) Call(name string, s *s.Stack) types.VType {
	var e types.VType = vnil.New()
	if n, np := t.native[name]; np {
		_, _, e = Run(&n, s, t)
	} else if f, fp := t.functions[name]; fp {
		e = f(s, t)
	} else {
		a, _ := t.aliases[name]
		return t.Call(a, s)
	}
	return e
}

func (t *Table) Has(name string) bool {
	_, n := t.native[name]
	_, f := t.functions[name]
	_, a := t.aliases[name]
	return n || f || a
}

func (t *Table) Define(name string, f Function) {
	t.functions[name] = f
}

func (t *Table) DefineNative(name string, ts *p.Tokens) {
	t.native[name] = *ts
}

func (t *Table) Alias(from, to string) {
	t.aliases[from] = to
}

func NewTable() *Table {
	tbl := new(Table)

	n := map[string]p.Tokens{}
	tbl.native = n

	f := map[string]Function{}
	tbl.functions = f

	a := map[string]string{}
	tbl.aliases = a

	return tbl
}
