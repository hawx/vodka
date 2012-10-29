// Package table implements a table of functions.
package table

import (
	p "github.com/hawx/vodka/parser"
	s "github.com/hawx/vodka/stack"

	"github.com/hawx/vodka/types"
)

// A function takes the current stack, the table and returns a value.
type Function func(*s.Stack, *Table) types.VType

// A table consists of natively defined functions, ie. those defined in vodka
// source code; functions defined in go code; and aliases.
type Table struct {
	Native    map[string]p.Tokens
	Function  map[string]Function
	Aliases   map[string]string
}

// String returns a formatted string displaying the names of all defined
// functions.
func (t *Table) String() string {
	s := "Native:\n"
	for i, ts := range t.Native {
		s += "  " + i + " " + ts.String() + "\n"
	}

	s += "\nFunctions:\n"
	for i, _ := range t.Function {
		s += "  " + i
	}

	return s + "\n"
}

// Defined returns a string of all function names joined by a single space.
func (t *Table) Defined() string {
	s := ""
	for n, _ := range t.Native {
		s += n + " "
	}
	for n, _ := range t.Function {
		s += n + " "
	}
	for n, _ := range t.Aliases {
		s += n + " "
	}
	return s
}

// Has returns true if the Table contains a function or alias with the name
// given.
func (t *Table) Has(name string) bool {
	_, n := t.Native[name]
	_, f := t.Function[name]
	_, a := t.Aliases[name]
	return n || f || a
}

// Define adds a new (go) function to the table with the name given.
func (t *Table) Define(name string, f Function) {
	t.Function[name] = f
}

// DefineNative adds a new (vodka) function to the table with the name given.
func (t *Table) DefineNative(name string, ts *p.Tokens) {
	t.Native[name] = *ts
}

// Alias defined a new alias so that when from is called, the function called to
// is run.
func (t *Table) Alias(from, to string) {
	t.Aliases[from] = to
}

// New returns a new, empty Table.
func New() *Table {
	tbl := new(Table)

	n := map[string]p.Tokens{}
	tbl.Native = n

	f := map[string]Function{}
	tbl.Function = f

	a := map[string]string{}
	tbl.Aliases = a

	return tbl
}
