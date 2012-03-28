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

// TOKENS ---------------------------------------------

type VTokens struct {
	value string
}

func (v *VTokens) String() string {
	return "[" + v.value + "]"
}

func (v *VTokens) Value() interface{} {
	return Parse(v.value)
}

func NewVTokens(s string) *VTokens {
	r := new(VTokens)
	r.value = s
	return r
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

		"string": func(s *Stack, t *Table) VType {
			v := s.popString()
			s.push(v)
			return v
		},

		// Basic I/O

		"print": func(s *Stack, t *Table) VType {
			v := s.popString()
			fmt.Println(v.String())
			return v
		},

		// Stack operations

		"pop": func(s *Stack, t *Table) VType {
			v := s.pop()
			return v
		},
		"size": func(s *Stack, t *Table) VType {
			v := NewVIntegerInt(s.size())
			s.push(v)
			return v
		},
		"dup": func(s *Stack, t *Table) VType {
			v := s.top()
			s.push(v)
			return v
		},
		"swap": func(s *Stack, t *Table) VType {
			a := s.pop()
			b := s.pop()
			s.push(b)
			s.push(a)
			return a
		},
		"swapp": func(s *Stack, t *Table) VType {
			a := s.pop()
			b := s.pop()
			c := s.pop()
			s.push(b)
			s.push(a)
			s.push(c)
			return c
		},
		"drop": func(s *Stack, t *Table) VType {
			s.clear()
			return VNil()
		},

		// Arithmetic operations

		"add": func(s *Stack, t *Table) VType {
			add := NewVIntegerInt(s.pop().Value().(int) + s.pop().Value().(int))
			s.push(add)
			return add
		},
		"prod": func(s *Stack, t *Table) VType {
			prod := NewVIntegerInt(s.pop().Value().(int) * s.pop().Value().(int))
			s.push(prod)
			return prod
		},
		"sub": func(s *Stack, t *Table) VType {
			sub := NewVIntegerInt(s.pop().Value().(int) - s.pop().Value().(int))
			s.push(sub)
			return sub
		},
		"div": func(s *Stack, t *Table) VType {
			div := NewVIntegerInt(s.pop().Value().(int) / s.pop().Value().(int))
			s.push(div)
			return div
		},

		// Logical

		"true": func(s *Stack, t *Table) VType {
			s.push(VTrue())
			return VNil()
		},
		"false": func(s *Stack, t *Table) VType {
			s.push(VFalse())
			return VNil()
		},
		"nil": func(s *Stack, t *Table) VType {
			s.push(VNil())
			return VNil()
		},

		"or": func(s *Stack, t *Table) VType {
			a := s.pop().Value().(bool)
			b := s.pop().Value().(bool)
			val := VFalse()
			if a || b {
				val = VTrue()
			}
			s.push(val)
			return VNil()
		},
		"and": func(s *Stack, t *Table) VType {
			a := s.pop().Value().(bool)
			b := s.pop().Value().(bool)
			val := VFalse()
			if a && b {
				val = VTrue()
			}
			s.push(val)
			return VNil()
		},
		"not": func(s *Stack, t *Table) VType {
			a := s.pop().Value().(bool)
			val := VTrue()
			if a {
				val = VFalse()
			}
			s.push(val)
			return VNil()
		},

		"eq?": func(s *Stack, t *Table) VType {
			val := VFalse()
			if s.pop().Value() == s.pop().Value() {
				val = VTrue()
			}
			s.push(val)
			return val
		},
		// "gt?": func(s *Stack, t *Table) VType {
		// 	val := VFalse()
		// 	if s.pop().Value() > s.pop().Value() {
		// 		val = VTrue()
		// 	}
		// 	s.push(val)
		// 	return val
		// },
		// "lt?": func(s *Stack, t *Table) VType {
		// 	val := VFalse()
		// 	if s.pop().Value() < s.pop().Value() {
		// 		val = VTrue()
		// 	}
		// 	s.push(val)
		// 	return val
		// },

		// Control flow

		"if": func(s *Stack, t *Table) VType {
			o := s.pop().Value().(*Tokens)
			cond := s.pop()
			if cond.Value().(bool) {
				s, t, _ = Run(o, s, t)
			}
			return VNil()
		},
		"unless": func(s *Stack, t *Table) VType {
			o := s.pop().Value().(*Tokens)
			cond := s.pop().Value().(bool)
			if cond {
				s, t, _ = Run(o, s, t)
			}
			return VNil()
		},
		"if-else": func(s *Stack, t *Table) VType {
			a := s.pop().Value().(*Tokens)
			b := s.pop().Value().(*Tokens)
			cond := s.pop().Value().(bool)
			if cond {
				s, t, _ = Run(b, s, t)
			} else {
				s, t, _ = Run(a, s, t)
			}
			return VNil()
		},
		"else-if": func(s *Stack, t *Table) VType {
			b := s.pop().Value().(*Tokens)
			a := s.pop().Value().(*Tokens)
			cond := s.pop().Value().(bool)
			if cond {
				s, t, _ = Run(b, s, t)
			} else {
				s, t, _ = Run(a, s, t)
			}
			return VNil()
		},
		"call": func(s *Stack, t *Table) VType {
			a := s.pop().Value().(*Tokens)
			Run(a, s, t)
			return VNil()
		},
	}

	tbl.functions = t

	tbl.alias("+", "add")
	tbl.alias("*", "prod")
	tbl.alias("-", "sub")
	tbl.alias("/", "div")

	return tbl
}


func BareEval(code string) {
	tokens := Parse(code)
	Run(tokens, NewStack(), BootedTable())
}

func Eval(code string, stk *Stack, tbl *Table) (s *Stack, t *Table, v VType) {
	tokens := Parse(code)
	return Run(tokens, stk, tbl)
}

func Run(tokens *Tokens, stk *Stack, tbl *Table) (s *Stack, t *Table, v VType) {
	var val VType = VNil()

	for _, tok := range *tokens {
		switch tok.key {
		case "str":
			stk.push(NewVInteger(tok.val))

		case "int":
			stk.push(NewVInteger(tok.val))

		case "stm":
			stk.push(NewVTokens(tok.val))

		case "fun":
			if tbl.has(tok.val) {
				f := tbl.get(tok.val)
				val = f(stk, tbl)
			} else {
				// no function error!
			}

		default:
			// problem?

		}
	}

	// If no value has been set show stack
	if val.Value() == nil {
		val = VNil()
	}

	return stk, tbl, val
}
