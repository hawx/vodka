// Package table implements a table of functions.
package table

import (
	"sort"

	p "hawx.me/code/vodka/parser"
	s "hawx.me/code/vodka/stack"
	"hawx.me/code/vodka/types"
)

type function func(*s.Stack, *Table) types.VType

// A function takes the current stack, the table and returns a value.
type Function struct {
	Apply function
}

// A table consists of natively defined functions, ie. those defined in vodka
// source code; functions defined in go code; and aliases.
type Table struct {
	Native   map[string]p.Tokens
	Function map[string]Function
	Aliases  map[string]string
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
func (t *Table) Defined() []string {
	s := []string{}

	for name, _ := range t.Native {
		s = append(s, name)
	}
	for name, _ := range t.Function {
		s = append(s, name)
	}
	for name, _ := range t.Aliases {
		s = append(s, name)
	}

	sort.Strings(s)
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
func (t *Table) Define(name string, f function) {
	t.Function[name] = Function{f}
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
	return &Table{
		Native:   map[string]p.Tokens{},
		Function: map[string]Function{},
		Aliases:  map[string]string{},
	}
}
