package main

import (
	"fmt"
	"strconv"
)

// TYPES --------------------------------------------

type VType interface {
	String()  string
	Value()   interface{}
}

// SPECIALS ---------------------------------------------

type VNilType struct { }

func (v *VNilType) String() string {
	return "nil"
}

func (v *VNilType) Value() interface{} {
	return nil
}

func VNil() *VNilType {
	r := new(VNilType)
	return r
}

// BOOLEAN ---------------------------------------------

type VBoolean struct {
	value bool
}

func (v *VBoolean) String() string {
	if v.value {
		return "true"
	}
	return "false"
}

func (v *VBoolean) Value() interface{} {
	return v.value
}

func VTrue() *VBoolean {
	b := new(VBoolean)
	b.value = true
	return b
}

func VFalse() *VBoolean {
	b := new(VBoolean)
	b.value = false
	return b
}

// STRING ---------------------------------------------

type VString struct {
	value string
}

func (v *VString) String() string {
	return v.value
}

func (v *VString) Value() interface{} {
	return v.value
}

func NewVString(s string) *VString {
	r := new(VString)
	r.value = s
	return r
}

// Integer ------------------------------------------------

type VInteger struct {
	value int
}

func (v *VInteger) String() string {
	return strconv.Itob(v.value, 10)
}

func (v *VInteger) Value() interface{} {
	return v.value
}

func NewVInteger(s string) *VInteger {
	r := new(VInteger)
	r.value, _ = strconv.Atoi(s)
	return r
}

func NewVIntegerInt(i int) *VInteger {
	r := new(VInteger)
	r.value = i
	return r
}


func BootedTable() *Table {
	tbl := NewTable()

	t := map[string] Function {

		// Type conversion

		"string": func(s *Stack) VType {
			v := s.popString()
			s.push(v)
			return v
		},

		// Basic I/O

		"print": func(s *Stack) VType {
			v := s.popString()
			fmt.Println()
			return v
		},

		// Stack operations

		"pop": func(s *Stack) VType {
			v := s.pop()
			return v
		},
		"size": func(s *Stack) VType {
			v := NewVIntegerInt(s.size())
			s.push(v)
			return v
		},
		"dup": func(s *Stack) VType {
			v := s.top()
			s.push(v)
			return v
		},
		"swap": func(s *Stack) VType {
			a := s.pop()
			b := s.pop()
			s.push(b)
			s.push(a)
			return a
		},
		"swapp": func(s *Stack) VType {
			a := s.pop()
			b := s.pop()
			c := s.pop()
			s.push(b)
			s.push(a)
			s.push(c)
			return c
		},
		"drop": func(s *Stack) VType {
			s.clear()
			return VNil()
		},

		// Arithmetic operations

		"add": func(s *Stack) VType {
			add := NewVIntegerInt(s.pop().Value().(int) + s.pop().Value().(int))
			s.push(add)
			return add
		},
		"prod": func(s *Stack) VType {
			prod := NewVIntegerInt(s.pop().Value().(int) * s.pop().Value().(int))
			s.push(prod)
			return prod
		},
		"sub": func(s *Stack) VType {
			sub := NewVIntegerInt(s.pop().Value().(int) - s.pop().Value().(int))
			s.push(sub)
			return sub
		},
		"div": func(s *Stack) VType {
			div := NewVIntegerInt(s.pop().Value().(int) / s.pop().Value().(int))
			s.push(div)
			return div
		},

		// Logical

		"true": func(s *Stack) VType {
			s.push(VTrue())
			return VNil()
		},
		"false": func(s *Stack) VType {
			s.push(VFalse())
			return VNil()
		},
		"nil": func(s *Stack) VType {
			s.push(VNil())
			return VNil()
		},

		"or": func(s *Stack) VType {
			a := s.pop().Value().(bool)
			b := s.pop().Value().(bool)
			val := VFalse()
			if a || b {
				val = VTrue()
			}
			s.push(val)
			return VNil()
		},
		"and": func(s *Stack) VType {
			a := s.pop().Value().(bool)
			b := s.pop().Value().(bool)
			val := VFalse()
			if a && b {
				val = VTrue()
			}
			s.push(val)
			return VNil()
		},
		"not": func(s *Stack) VType {
			a := s.pop().Value().(bool)
			val := VTrue()
			if a {
				val = VFalse()
			}
			s.push(val)
			return VNil()
		},

	}

	tbl.functions = t

	tbl.alias("+", "add")
	tbl.alias("*", "prod")

	return tbl
}


func Eval(code string, stk *Stack, tbl *Table) (s *Stack, t *Table, v string) {
	tokens := Parse(code)
	return Run(tokens, stk, tbl)
}

func Run(tokens *Tokens, stk *Stack, tbl *Table) (s *Stack, t *Table, v string) {
	var val VType = VNil()

	for _, tok := range *tokens {
		switch tok.key {
		case "str":
			stk.push(NewVInteger(tok.val))

		case "int":
			stk.push(NewVInteger(tok.val))

		case "stm":
			// add statements to stack

		case "fun":
			if tbl.has(tok.val) {
				f := tbl.get(tok.val)
				val = f(stk)
			} else {
				// no function error!
			}

		default:
			// problem?

		}
	}

	// If no value has been set show stack
	if val.Value() == nil {
		val = NewVString(stk.String())
	}

	return stk, tbl, val.String()
}
