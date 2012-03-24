package main

import (
	"fmt"
	"strconv"
)

type VType interface {
	String()  string
}

type VString struct {
	value string
}

func (v *VString) String() string {
	return v.value
}

type VInteger struct {
	value int
}

func (v *VInteger) String() string {
	return strconv.Itob(v.value, 10)
}

func BootedTable() *Table {
	tbl := NewTable()

	t := map[string] Function {

		// Type conversion

		"string": func(s *Stack) *Stack {
			s.push(s.popString())
			return s
		},

		// Basic I/O

		"print": func(s *Stack) *Stack {
			fmt.Println(s.popString())
			return s
		},

		// Stack operations

		"pop": func(s *Stack) *Stack {
			s.pop()
			return s
		},
		"size": func(s *Stack) *Stack {
			s.push(s.size())
			return s
		},
		"dup": func(s *Stack) *Stack {
			s.push(s.top())
			return s
		},
		"swap": func(s *Stack) *Stack {
			a := s.pop()
			b := s.pop()
			s.push(b)
			s.push(a)
			return s
		},
		"drop": func(s *Stack) *Stack {
			s.clear()
			return s
		},

		// Arithmetic operations

		"add": func(s *Stack) *Stack {
			add := s.pop().(int) + s.pop().(int)
			s.push(add)
			return s
		},
		"prod": func(s *Stack) *Stack {
			prod := s.pop().(int) * s.pop().(int)
			s.push(prod)
			return s
		},

	}

	tbl.functions = t

	tbl.alias("+", "add")
	tbl.alias("*", "prod")

	return tbl
}


func Eval(code string, stk *Stack, tbl *Table) (s *Stack, t *Table) {
	tokens := Parse(code)
	return Run(tokens, stk, tbl)
}

func Run(tokens *Tokens, stk *Stack, tbl *Table) (s *Stack, t *Table) {

	for _, tok := range *tokens {
		switch tok.key {
		case "str":
			stk.push(tok.val)

		case "int":
			i, _ := strconv.Atoi(tok.val)
			stk.push(i)

		case "stm":
			// add statements to stack

		case "fun":
			if tbl.has(tok.val) {
				f := tbl.get(tok.val)
				stk = f(stk)
			} else {
				// no function error!
			}

		default:
			// problem?

		}
	}

	fmt.Println(stk.String())

	return stk, tbl
}
