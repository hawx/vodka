package main

import (
	"fmt"
	"io/ioutil"
)

func BootedTable() *Table {
	tbl := NewTable()

	t := map[string]Function{

		"eval": func(s *Stack, t *Table) VType {
			str := s.Pop().Value().(string)
			_, _, v := Eval(str, s, t)
			return v
		},
		"alias": func(s *Stack, t *Table) VType {
			from := s.Pop().Value().(string)
			to := s.Pop().Value().(string)
			t.Alias(from, to)
			return VNil()
		},
		"define": func(s *Stack, t *Table) VType {
			stms := s.Pop().Value().(*Tokens)
			name := s.Pop().Value().(string)
			t.DefineNative(name, stms)
			return VNil()
		},
		"defined": func(s *Stack, t *Table) VType {
			v := NewVString(t.Defined())
			s.Push(v)
			return VNil()
		},
		"type": func(s *Stack, t *Table) VType {
			v := NewVString(s.Pop().Type())
			s.Push(v)
			return VNil()
		},

		// Basic I/O

		"print": func(s *Stack, t *Table) VType {
			v := s.Pop().Value().(string)
			fmt.Println(v)
			return VNil()
		},
		"p": func(s *Stack, t *Table) VType {
			v := s.PopString().String()
			fmt.Println(v[1 : len(v)-1])
			return VNil()
		},
		"read": func(s *Stack, t *Table) VType {
			contents, _ := ioutil.ReadFile(s.Pop().Value().(string))
			str := NewVString(string(contents))
			s.Push(str)
			return VNil()
		},

		// Stack operations

		"pop": func(s *Stack, t *Table) VType {
			v := s.Pop()
			return v
		},
		"size": func(s *Stack, t *Table) VType {
			v := NewVIntegerInt(s.Size())
			s.Push(v)
			return VNil()
		},
		"dup": func(s *Stack, t *Table) VType {
			v := s.Top()
			s.Push(v)
			return VNil()
		},
		"swap": func(s *Stack, t *Table) VType {
			a := s.Pop()
			b := s.Pop()
			s.Push(a)
			s.Push(b)
			return VNil()
		},
		"drop": func(s *Stack, t *Table) VType {
			s.Clear()
			return VNil()
		},
		"compose": func(s *Stack, t *Table) VType {
			a := s.Pop()
			b := s.Pop()
			c := NewVBlock(b.(*VBlock).value + " " + a.(*VBlock).value)
			s.Push(c)
			return VNil()
		},
		"wrap": func(s *Stack, t *Table) VType {
			b := s.Pop()
			r := NewVBlock("[" + b.(*VBlock).value + "]")
			s.Push(r)
			return VNil()
		},

		// Arithmetic operations

		"add": func(s *Stack, t *Table) VType {
			add := NewVIntegerInt(s.Pop().Value().(int) + s.Pop().Value().(int))
			s.Push(add)
			return VNil()
		},
		"mult": func(s *Stack, t *Table) VType {
			mult := NewVIntegerInt(s.Pop().Value().(int) * s.Pop().Value().(int))
			s.Push(mult)
			return VNil()
		},
		"sub": func(s *Stack, t *Table) VType {
			a := s.Pop().Value().(int)
			b := s.Pop().Value().(int)
			sub := NewVIntegerInt(b - a)
			s.Push(sub)
			return VNil()
		},
		"div": func(s *Stack, t *Table) VType {
			a := s.Pop().Value().(int)
			b := s.Pop().Value().(int)
			div := NewVIntegerInt(b / a)
			s.Push(div)
			return VNil()
		},
		"neg": func(s *Stack, t *Table) VType {
			val := NewVIntegerInt(-s.Pop().Value().(int))
			s.Push(val)
			return VNil()
		},

		// Logical

		"true": func(s *Stack, t *Table) VType {
			s.Push(VTrue())
			return VNil()
		},
		"false": func(s *Stack, t *Table) VType {
			s.Push(VFalse())
			return VNil()
		},
		"nil": func(s *Stack, t *Table) VType {
			s.Push(VNil())
			return VNil()
		},

		"or": func(s *Stack, t *Table) VType {
			a := s.Pop().Value().(bool)
			b := s.Pop().Value().(bool)
			val := VFalse()
			if a || b {
				val = VTrue()
			}
			s.Push(val)
			return VNil()
		},
		"and": func(s *Stack, t *Table) VType {
			a := s.Pop().Value().(bool)
			b := s.Pop().Value().(bool)
			val := VFalse()
			if a && b {
				val = VTrue()
			}
			s.Push(val)
			return VNil()
		},

		"compare": func(s *Stack, t *Table) VType {
			a := s.Pop()
			b := s.Pop()
			val := NewVIntegerInt(a.CompareWith(b))
			s.Push(val)
			return VNil()
		},
		"eq?": func(s *Stack, t *Table) VType {
			a := s.Pop()
			b := s.Pop()
			val := NewVBoolean(a.CompareWith(b) == 0)
			s.Push(val)
			return VNil()
		},

		// Control flow

		"if-else": func(s *Stack, t *Table) VType {
			a := s.Pop().Value().(*Tokens)
			b := s.Pop().Value().(*Tokens)
			cond := s.Pop().Value().(bool)
			if cond {
				s, t, _ = Run(a, s, t)
			} else {
				s, t, _ = Run(b, s, t)
			}
			return VNil()
		},
		"call": func(s *Stack, t *Table) VType {
			val := s.Top().Value()
			switch val.(type) {
			case *Tokens:
				s.Pop()
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
			save := s.Pop()
			tokens := new(Tokens)
			*tokens = append(*tokens, NewToken("fun", "call"))
			Run(tokens, s, t)
			s.Push(save)
			return VNil()
		},
		"without2": func(s *Stack, t *Table) VType {
			save1 := s.Pop()
			save2 := s.Pop()
			tokens := new(Tokens)
			*tokens = append(*tokens, NewToken("fun", "call"))
			Run(tokens, s, t)
			s.Push(save2)
			s.Push(save1)
			return VNil()
		},

		// Strings

		"string": func(s *Stack, t *Table) VType {
			v := s.PopString()
			s.Push(v)
			return VNil()
		},
		"concat": func(s *Stack, t *Table) VType {
			a := s.Pop().Value().(string)
			b := s.Pop().Value().(string)
			c := NewVString(b + a)
			s.Push(c)
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
			stk.Push(NewVString(tok.val))

		case "int":
			stk.Push(NewVInteger(tok.val))

		case "stm":
			stk.Push(NewVBlock(tok.val))

		case "fun":
			if tbl.Has(tok.val) {
				val = tbl.Call(tok.val, stk)
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
