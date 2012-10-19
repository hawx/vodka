package main

import (
	"fmt"
	"io/ioutil"
	p "./parser"
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
			stms := s.Pop().Value().(*p.Tokens)
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


		// Types

		"integer": func(s *Stack, t *Table) VType {
			v := s.Pop().Value().(string)
			s.Push(NewVInteger(v))
			return VNil()
		},
		"string": func(s *Stack, t *Table) VType {
			v := s.PopString()
			s.Push(v)
			return VNil()
		},
		"list": func(s *Stack, t *Table) VType {
			v := s.Pop()
			list := make([]VType, 1)
			list[0] = v

			s.Push(NewVListList(list))

			return VNil()
		},
		"range": func(s *Stack, t *Table) VType {
			end   := s.Pop().Value().(int)
			start := s.Pop().Value().(int)

			list := make([]VType, end - start + 1)
			for i := start; i <= end; i++ {
				list[i-start] = NewVIntegerInt(i)
			}
			s.Push(NewVListList(list))

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
			r := NewVBlock(b.String())
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
			val := NewVIntegerInt(a.Compare(b))
			s.Push(val)
			return VNil()
		},
		"eq?": func(s *Stack, t *Table) VType {
			a := s.Pop()
			b := s.Pop()
			val := NewVBoolean(a.Compare(b) == 0)
			s.Push(val)
			return VNil()
		},

		// Control flow

		"if-else": func(s *Stack, t *Table) VType {
			a := s.Pop().Value().(*p.Tokens)
			b := s.Pop().Value().(*p.Tokens)
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
			case *p.Tokens:
				s.Pop()
				Run(val.(*p.Tokens), s, t)
			case *VBlock:
				toks := new(p.Tokens)
				*toks = append(*toks, p.Token{"fun", "call"})
				Run(toks, s, t)
			default:
				println("Unexpected type")
			}
			return VNil()
		},
		"without": func(s *Stack, t *Table) VType {
			save := s.Pop()
			tokens := new(p.Tokens)
			*tokens = append(*tokens, p.Token{"fun", "call"})
			Run(tokens, s, t)
			s.Push(save)
			return VNil()
		},
		"without2": func(s *Stack, t *Table) VType {
			save1 := s.Pop()
			save2 := s.Pop()
			tokens := new(p.Tokens)
			*tokens = append(*tokens, p.Token{"fun", "call"})
			Run(tokens, s, t)
			s.Push(save2)
			s.Push(save1)
			return VNil()
		},

		// Strings

		"concat": func(s *Stack, t *Table) VType {
			a := s.Pop().Value().(string)
			b := s.Pop().Value().(string)
			c := NewVString(b + a)
			s.Push(c)
			return VNil()
		},

		// Lists

		"head": func(s *Stack, t *Table) VType {
			h := s.Pop().Value().([]VType)[0]
			s.Push(h)
			return VNil()
		},
		"tail": func(s *Stack, t *Table) VType {
			v := s.Pop().Value().([]VType)[1:]
			s.Push(NewVListList(v))
			return VNil()
		},
		"cons": func(s *Stack, t *Table) VType {
			v := s.Pop().(VType)
			l := s.Pop().Value().([]VType)
			l = append(l, v)
			s.Push(NewVListList(l))
			return VNil()
		},
		"append": func(s *Stack, t *Table) VType {
			a := s.Pop().Value().([]VType)
			b := s.Pop().Value().([]VType)

			// Ripped from http://golang.org/doc/effective_go.html#slices
			l := len(a)
			if l + len(b) > cap(a) {
				// Allocate double what's needed, for future growth.
				newSlice := make([]VType, (l + len(b)) * 2)
				copy(newSlice, a)
				a = newSlice
			}
			a = a[0:l+len(b)]
			for i, c := range b {
				a[l+i] = c
			}

			s.Push(NewVListList(a))
			return VNil()
		},
		"apply": func(s *Stack, t *Table) VType {
			f := s.Pop().Value().(*p.Tokens)
			l := s.Pop().Value().([]VType)

			stk := make(Stack, len(l))
			for i, o := range l {
				stk[i] = o.(VType)
			}

			newstk, _, v := Run(f, &stk, t)

			list := make([]VType, len(*newstk))
			for i, o := range *newstk {
				list[i] = o.(VType)
			}
			s.Push(NewVListList(list))

			return v
		},
		"reverse": func(s *Stack, t *Table) VType {
			l := s.Pop().Value().([]VType)

			for i, j := 0, len(l)-1; i < j; i, j = i+1, j-1 {
				l[i], l[j] = l[j], l[i]
			}

			s.Push(NewVListList(l))

			return VNil()
		},
	}

	tbl.functions = t

	contents, _ := ioutil.ReadFile("boot.vk")
	_, tbl, _ = Eval(string(contents), NewStack(), tbl)

	return tbl
}

func BareEval(code string) {
	tokens := p.Parse(code)
	Run(tokens, NewStack(), BootedTable())
}

func Eval(code string, stk *Stack, tbl *Table) (s *Stack, t *Table, v VType) {
	tokens := p.Parse(code)
	return Run(tokens, stk, tbl)
}

func Run(tokens *p.Tokens, stk *Stack, tbl *Table) (s *Stack, t *Table, v VType) {
	var val VType = VNil()

	for _, tok := range *tokens {
		switch tok.Key {
		case "str":
			stk.Push(NewVString(tok.Val))

		case "int":
			stk.Push(NewVInteger(tok.Val))

		case "list":
			// TODO: need to make sure you can't add to ^tbl^ from a list.
			sub := NewStack()
			sub, _, _ = Eval(tok.Val, sub, tbl)
			stk.Push(NewVList(sub))

		case "stm":
			stk.Push(NewVBlock(tok.Val))

		case "fun":
			if tbl.Has(tok.Val) {
				val = tbl.Call(tok.Val, stk)
			} else {
				println("Unknown function: '" + tok.Val + "'")
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
