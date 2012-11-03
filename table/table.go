// Package table implements a table of functions.
package table

import (
	"fmt"
	"strings"

	p "github.com/hawx/vodka/parser"
	s "github.com/hawx/vodka/stack"
	"github.com/hawx/vodka/types"
)

// A type signature for a vodka function.
type Signature struct {
	// The expected elements on the stack.
	In   []string
	// The output of the function to the stack.
	Out  []string
}

func ParseSignature(sig string) Signature {
	parts := strings.Split(sig, "->")
	if len(parts) != 2 {
		fmt.Println("Malformed signature: '", sig, "'.")
	}

	lhs := parts[0]; rhs := parts[1]

	left  := strings.Split(strings.TrimSpace(lhs), " ")
	right := strings.Split(strings.TrimSpace(rhs), " ")

	for i,_ := range left {
		left[i] = strings.TrimSpace(left[i])
	}

	if len(left) == 1 && left[0] == "" { left = []string{} }

	for i,_ := range right {
		right[i] = strings.TrimSpace(right[i])
	}

	if len(right) == 1 && right[0] == "" { right = []string{} }

	return Signature{left, right}
}

func (s Signature) String() string {
	return strings.Join(s.In, " ") + " -> " + strings.Join(s.Out, " ")
}

func (s Signature) Check(stk *s.Stack) bool {
	if len(s.In) > 0 && s.In[0] == "'A" { s.In = s.In[1:len(s.In)] }
	top := stk.Peek(len(s.In))

	if len(s.In) != len(top) { return false }

	for i,e := range top {
		if s.In[i][0:1] == "'" {
			// wildcard, ignore.
		} else if s.In[i] != e.Type() {
			return false
		}
	}

	return true
}

type function func(*s.Stack, *Table) types.VType

// A function takes the current stack, the table and returns a value.
type Function struct {
	Apply  function
	Sig    Signature
}

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
	func (t *Table) Define(name string, f function, sig string) {
	t.Function[name] = Function{f, ParseSignature(sig)}
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
