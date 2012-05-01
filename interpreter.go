package main

import (
	"fmt"
	"io/ioutil"
)

func BootedTable() *Table {
	tbl := NewTable()

	t := map[string]Function{

		"__document__": func(s *Stack, t *Table) VType {
			return VNil()
		},
		"eval": func(s *Stack, t *Table) VType {
			str := s.pop().Value().(string)
			_, _, v := Eval(str, s, t)
			return v
		},

		// Reflection

		"defined": func(s *Stack, t *Table) VType {
			v := NewVString(t.Defined())
			s.push(v)
			return VNil()
		},
		"type": func(s *Stack, t *Table) VType {
			v := NewVString(s.pop().Type())
			s.push(v)
			return VNil()
		},

		// Type conversion

		"string": func(s *Stack, t *Table) VType {
			v := s.popString()
			s.push(v)
			return VNil()
		},
		"concat": func(s *Stack, t *Table) VType {
			a := s.pop().Value().(string)
			b := s.pop().Value().(string)
			c := NewVString(b + a)
			s.push(c)
			return VNil()
		},

		// Basic I/O

		"print": func(s *Stack, t *Table) VType {
			v := s.pop().Value().(string)
			fmt.Println(v)
			return VNil()
		},
		"p": func(s *Stack, t *Table) VType {
			v := s.popString().String()
			fmt.Println(v[1 : len(v)-1])
			return VNil()
		},
		"read": func(s *Stack, t *Table) VType {
			contents, _ := ioutil.ReadFile(s.pop().Value().(string))
			str := NewVString(string(contents))
			s.push(str)
			return VNil()
		},

		// Stack operations

		"pop": func(s *Stack, t *Table) VType {
			v := s.pop()
			return v
		},
		"size": func(s *Stack, t *Table) VType {
			v := NewVIntegerInt(s.size())
			s.push(v)
			return VNil()
		},
		"dup": func(s *Stack, t *Table) VType {
			v := s.top()
			s.push(v)
			return VNil()
		},
		"swap": func(s *Stack, t *Table) VType {
			a := s.pop()
			b := s.pop()
			s.push(a)
			s.push(b)
			return VNil()
		},
		"swapp": func(s *Stack, t *Table) VType {
			a := s.pop()
			b := s.pop()
			c := s.pop()
			s.push(b)
			s.push(c)
			s.push(a)
			return VNil()
		},
		"drop": func(s *Stack, t *Table) VType {
			s.clear()
			return VNil()
		},
		"compose": func(s *Stack, t *Table) VType {
			a := s.pop()
			b := s.pop()
			c := NewVBlock(b.(*VBlock).value + " " + a.(*VBlock).value)
			s.push(c)
			return VNil()
		},
		"wrap": func(s *Stack, t *Table) VType {
			b := s.pop()
			r := NewVBlock("[" + b.(*VBlock).value + "]")
			s.push(r)
			return VNil()
		},

		// Arithmetic operations

		"add": func(s *Stack, t *Table) VType {
			add := NewVIntegerInt(s.pop().Value().(int) + s.pop().Value().(int))
			s.push(add)
			return VNil()
		},
		"mult": func(s *Stack, t *Table) VType {
			mult := NewVIntegerInt(s.pop().Value().(int) * s.pop().Value().(int))
			s.push(mult)
			return VNil()
		},
		"sub": func(s *Stack, t *Table) VType {
			a := s.pop().Value().(int)
			b := s.pop().Value().(int)
			sub := NewVIntegerInt(b - a)
			s.push(sub)
			return VNil()
		},
		"div": func(s *Stack, t *Table) VType {
			a := s.pop().Value().(int)
			b := s.pop().Value().(int)
			div := NewVIntegerInt(b / a)
			s.push(div)
			return VNil()
		},
		"neg": func(s *Stack, t *Table) VType {
			val := NewVIntegerInt(-s.pop().Value().(int))
			s.push(val)
			return VNil()
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

		"compare": func(s *Stack, t *Table) VType {
			a := s.pop()
			b := s.pop()
			val := NewVIntegerInt(a.CompareWith(b))
			s.push(val)
			return VNil()
		},
		"eq?": func(s *Stack, t *Table) VType {
			a := s.pop()
			b := s.pop()
			val := NewVBoolean(a.CompareWith(b) == 0)
			s.push(val)
			return VNil()
		},

		// Control flow

		"if-else": func(s *Stack, t *Table) VType {
			a := s.pop().Value().(*Tokens)
			b := s.pop().Value().(*Tokens)
			cond := s.pop().Value().(bool)
			if cond {
				s, t, _ = Run(a, s, t)
			} else {
				s, t, _ = Run(b, s, t)
			}
			return VNil()
		},
		"call": func(s *Stack, t *Table) VType {
			val := s.top().Value()
			switch val.(type) {
			case *Tokens:
				s.pop()
				Run(val.(*Tokens), s, t)
			case *VBlock:
				toks := new(Tokens)
				*toks = append(*toks, NewToken("fun", "call"))
				Run(toks, s, t)
			default:
				println("Unexpected type")
			}
			return VNil()
		},
		"without": func(s *Stack, t *Table) VType {
			save := s.pop()
			tokens := new(Tokens)
			*tokens = append(*tokens, NewToken("fun", "call"))
			Run(tokens, s, t)
			s.push(save)
			return VNil()
		},
		"without2": func(s *Stack, t *Table) VType {
			save1 := s.pop()
			save2 := s.pop()
			tokens := new(Tokens)
			*tokens = append(*tokens, NewToken("fun", "call"))
			Run(tokens, s, t)
			s.push(save2)
			s.push(save1)
			return VNil()
		},


		// Definition

		"alias": func(s *Stack, t *Table) VType {
			from := s.pop().Value().(string)
			to := s.pop().Value().(string)
			t.alias(from, to)
			return VNil()
		},
		"define": func(s *Stack, t *Table) VType {
			stms := s.pop().Value().(*Tokens)
			name := s.pop().Value().(string)
			t.defineNative(name, stms)
			return VNil()
		},
	}

	tbl.functions = t

	contents, _ := ioutil.ReadFile("boot.vk")
	_, tbl, _ = Eval(string(contents), NewStack(), tbl)

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
			stk.push(NewVString(tok.val))

		case "int":
			stk.push(NewVInteger(tok.val))

		case "stm":
			stk.push(NewVBlock(tok.val))

		case "fun":
			if tbl.has(tok.val) {
				val = tbl.call(tok.val, stk)
			} else {
				println("Unknown function: '" + tok.val + "'")
			}

		default:
			println("Unknown token: " + tok.String())
		}
	}

	// If no value has been set show stack
	if val.Value() == nil {
		val = VNil()
	}

	return stk, tbl, val
}
