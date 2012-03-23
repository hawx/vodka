package main

import (
	"fmt"
	"strconv"
)


func BootedTable() *Table {
	tbl := NewTable()

	tbl.define("add", func(s *Stack) *Stack {
		add := s.pop().(int) + s.pop().(int)
		s.push(add)
		return s
	})

	tbl.define("prod", func(s *Stack) *Stack {
		prod := s.pop().(int) * s.pop().(int)
		s.push(prod)
		return s
	})

	tbl.alias("+", "add")
	tbl.alias("*", "prod")

	return tbl
}


func BareEval(code string) (s *Stack, t *Table) {
	tokens := Parse(code)
	return BareRun(tokens)
}

func Eval(code string, stk *Stack, tbl *Table) (s *Stack, t *Table) {
	tokens := Parse(code)
	return Run(tokens, stk, tbl)
}

func BareRun(tokens *Tokens) (s *Stack, t *Table) {
	stk := NewStack()
	tbl := BootedTable()
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
