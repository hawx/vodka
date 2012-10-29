package interpreter

import (
	"fmt"
	"io/ioutil"
	p "github.com/hawx/vodka/parser"

	"github.com/hawx/vodka/table"
	"github.com/hawx/vodka/stack"

	"github.com/hawx/vodka/types"
	"github.com/hawx/vodka/types/vblock"
	"github.com/hawx/vodka/types/vboolean"
	"github.com/hawx/vodka/types/vnil"
	"github.com/hawx/vodka/types/vstring"
	"github.com/hawx/vodka/types/vinteger"
	"github.com/hawx/vodka/types/vlist"
)

func BootedTable() *table.Table {
	tbl := table.New()

	t := map[string]table.Function{

		"eval": func(s *stack.Stack, t *table.Table) types.VType {
			str := s.Pop().Value().(string)
			_, _, v := Eval(str, s, t)
			return v
		},
		"alias": func(s *stack.Stack, t *table.Table) types.VType {
			from := s.Pop().Value().(string)
			to := s.Pop().Value().(string)
			t.Alias(from, to)
			return vnil.New()
		},
		"define": func(s *stack.Stack, t *table.Table) types.VType {
			stms := s.Pop().Value().(*p.Tokens)
			name := s.Pop().Value().(string)
			t.DefineNative(name, stms)
			return vnil.New()
		},
		"defined": func(s *stack.Stack, t *table.Table) types.VType {
			v := vstring.New(t.Defined())
			s.Push(v)
			return vnil.New()
		},
		"type": func(s *stack.Stack, t *table.Table) types.VType {
			v := vstring.New(s.Pop().Type())
			s.Push(v)
			return vnil.New()
		},


		// Types

		"integer": func(s *stack.Stack, t *table.Table) types.VType {
			v := s.Pop().Value().(string)
			s.Push(vinteger.New(v))
			return vnil.New()
		},
		"string": func(s *stack.Stack, t *table.Table) types.VType {
			v := s.PopString()
			s.Push(v)
			return vnil.New()
		},
		"list": func(s *stack.Stack, t *table.Table) types.VType {
			v := s.Pop()
			list := make([]types.VType, 1)
			list[0] = v

			s.Push(vlist.NewFromList(list))

			return vnil.New()
		},
		"range": func(s *stack.Stack, t *table.Table) types.VType {
			end   := s.Pop().Value().(int)
			start := s.Pop().Value().(int)

			list := make([]types.VType, end - start + 1)
			for i := start; i <= end; i++ {
				list[i-start] = vinteger.NewFromInt(i)
			}
			s.Push(vlist.NewFromList(list))

			return vnil.New()
		},


		// Basic I/O

		"print": func(s *stack.Stack, t *table.Table) types.VType {
			v := s.Pop().Value().(string)
			fmt.Println(v)
			return vnil.New()
		},
		"p": func(s *stack.Stack, t *table.Table) types.VType {
			v := s.PopString().String()
			fmt.Println(v[1 : len(v)-1])
			return vnil.New()
		},
		"read": func(s *stack.Stack, t *table.Table) types.VType {
			contents, _ := ioutil.ReadFile(s.Pop().Value().(string))
			str := vstring.New(string(contents))
			s.Push(str)
			return vnil.New()
		},

		// s.Stack operations

		"pop": func(s *stack.Stack, t *table.Table) types.VType {
			v := s.Pop()
			return v
		},
		"size": func(s *stack.Stack, t *table.Table) types.VType {
			v := vinteger.NewFromInt(s.Size())
			s.Push(v)
			return vnil.New()
		},
		"dup": func(s *stack.Stack, t *table.Table) types.VType {
			v := s.Top()
			s.Push(v)
			return vnil.New()
		},
		"swap": func(s *stack.Stack, t *table.Table) types.VType {
			a := s.Pop()
			b := s.Pop()
			s.Push(a)
			s.Push(b)
			return vnil.New()
		},
		"drop": func(s *stack.Stack, t *table.Table) types.VType {
			s.Clear()
			return vnil.New()
		},
		"compose": func(s *stack.Stack, t *table.Table) types.VType {
			a := s.Pop()
			b := s.Pop()
			c := vblock.New(b.(*vblock.VBlock).BareValue() + " " + a.(*vblock.VBlock).BareValue())
			s.Push(c)
			return vnil.New()
		},
		"wrap": func(s *stack.Stack, t *table.Table) types.VType {
			b := s.Pop()
			r := vblock.New(b.String())
			s.Push(r)
			return vnil.New()
		},

		// Arithmetic operations

		"add": func(s *stack.Stack, t *table.Table) types.VType {
			add := vinteger.NewFromInt(s.Pop().Value().(int) + s.Pop().Value().(int))
			s.Push(add)
			return vnil.New()
		},
		"mult": func(s *stack.Stack, t *table.Table) types.VType {
			mult := vinteger.NewFromInt(s.Pop().Value().(int) * s.Pop().Value().(int))
			s.Push(mult)
			return vnil.New()
		},
		"sub": func(s *stack.Stack, t *table.Table) types.VType {
			a := s.Pop().Value().(int)
			b := s.Pop().Value().(int)
			sub := vinteger.NewFromInt(b - a)
			s.Push(sub)
			return vnil.New()
		},
		"div": func(s *stack.Stack, t *table.Table) types.VType {
			a := s.Pop().Value().(int)
			b := s.Pop().Value().(int)
			div := vinteger.NewFromInt(b / a)
			s.Push(div)
			return vnil.New()
		},
		"neg": func(s *stack.Stack, t *table.Table) types.VType {
			val := vinteger.NewFromInt(-s.Pop().Value().(int))
			s.Push(val)
			return vnil.New()
		},

		// Logical

		"true": func(s *stack.Stack, t *table.Table) types.VType {
			s.Push(vboolean.True())
			return vnil.New()
		},
		"false": func(s *stack.Stack, t *table.Table) types.VType {
			s.Push(vboolean.False())
			return vnil.New()
		},
		"nil": func(s *stack.Stack, t *table.Table) types.VType {
			s.Push(vnil.New())
			return vnil.New()
		},

		"or": func(s *stack.Stack, t *table.Table) types.VType {
			a := s.Pop().Value().(bool)
			b := s.Pop().Value().(bool)
			val := vboolean.False()
			if a || b {
				val = vboolean.True()
			}
			s.Push(val)
			return vnil.New()
		},
		"and": func(s *stack.Stack, t *table.Table) types.VType {
			a := s.Pop().Value().(bool)
			b := s.Pop().Value().(bool)
			val := vboolean.False()
			if a && b {
				val = vboolean.True()
			}
			s.Push(val)
			return vnil.New()
		},

		"compare": func(s *stack.Stack, t *table.Table) types.VType {
			a := s.Pop()
			b := s.Pop()
			val := vinteger.NewFromInt(a.Compare(b))
			s.Push(val)
			return vnil.New()
		},
		"eq?": func(s *stack.Stack, t *table.Table) types.VType {
			a := s.Pop()
			b := s.Pop()
			val := vboolean.New(a.Compare(b) == 0)
			s.Push(val)
			return vnil.New()
		},

		// Control flow

		"if-else": func(s *stack.Stack, t *table.Table) types.VType {
			a := s.Pop().Value().(*p.Tokens)
			b := s.Pop().Value().(*p.Tokens)
			cond := s.Pop().Value().(bool)
			if cond {
				s, t, _ = Run(a, s, t)
			} else {
				s, t, _ = Run(b, s, t)
			}
			return vnil.New()
		},
		"call": func(s *stack.Stack, t *table.Table) types.VType {
			val := s.Top().Value()
			switch val.(type) {
			case *p.Tokens:
				s.Pop()
				Run(val.(*p.Tokens), s, t)
			case *vblock.VBlock:
				toks := new(p.Tokens)
				*toks = append(*toks, p.Token{"fun", "call"})
				Run(toks, s, t)
			default:
				println("Unexpected type")
			}
			return vnil.New()
		},
		"without": func(s *stack.Stack, t *table.Table) types.VType {
			save := s.Pop()
			tokens := new(p.Tokens)
			*tokens = append(*tokens, p.Token{"fun", "call"})
			Run(tokens, s, t)
			s.Push(save)
			return vnil.New()
		},
		"without2": func(s *stack.Stack, t *table.Table) types.VType {
			save1 := s.Pop()
			save2 := s.Pop()
			tokens := new(p.Tokens)
			*tokens = append(*tokens, p.Token{"fun", "call"})
			Run(tokens, s, t)
			s.Push(save2)
			s.Push(save1)
			return vnil.New()
		},

		// Strings

		"concat": func(s *stack.Stack, t *table.Table) types.VType {
			a := s.Pop().Value().(string)
			b := s.Pop().Value().(string)
			c := vstring.New(b + a)
			s.Push(c)
			return vnil.New()
		},

		// Lists

		"head": func(s *stack.Stack, t *table.Table) types.VType {
			h := s.Pop().Value().([]types.VType)[0]
			s.Push(h)
			return vnil.New()
		},
		"tail": func(s *stack.Stack, t *table.Table) types.VType {
			v := s.Pop().Value().([]types.VType)[1:]
			s.Push(vlist.NewFromList(v))
			return vnil.New()
		},
		"cons": func(s *stack.Stack, t *table.Table) types.VType {
			v := s.Pop().(types.VType)
			l := s.Pop().Value().([]types.VType)
			l = append(l, v)
			s.Push(vlist.NewFromList(l))
			return vnil.New()
		},
		"append": func(s *stack.Stack, t *table.Table) types.VType {
			a := s.Pop().Value().([]types.VType)
			b := s.Pop().Value().([]types.VType)

			// Ripped from http://golang.org/doc/effective_go.html#slices
			l := len(a)
			if l + len(b) > cap(a) {
				// Allocate double what's needed, for future growth.
				newSlice := make([]types.VType, (l + len(b)) * 2)
				copy(newSlice, a)
				a = newSlice
			}
			a = a[0:l+len(b)]
			for i, c := range b {
				a[l+i] = c
			}

			s.Push(vlist.NewFromList(a))
			return vnil.New()
		},
		"apply": func(s *stack.Stack, t *table.Table) types.VType {
			f := s.Pop().Value().(*p.Tokens)
			l := s.Pop().Value().([]types.VType)

			stk := make(stack.Stack, len(l))
			for i, o := range l {
				stk[i] = o.(types.VType)
			}

			newstk, _, v := Run(f, &stk, t)

			list := make([]types.VType, len(*newstk))
			for i, o := range *newstk {
				list[i] = o.(types.VType)
			}
			s.Push(vlist.NewFromList(list))

			return v
		},
		"reverse": func(s *stack.Stack, t *table.Table) types.VType {
			l := s.Pop().Value().([]types.VType)

			for i, j := 0, len(l)-1; i < j; i, j = i+1, j-1 {
				l[i], l[j] = l[j], l[i]
			}

			s.Push(vlist.NewFromList(l))

			return vnil.New()
		},
	}

	tbl.Function = t

	contents, _ := ioutil.ReadFile("boot.vk")
	_, tbl, _ = Eval(string(contents), stack.New(), tbl)

	return tbl
}

func Eval(code string, stk *stack.Stack, tbl *table.Table) (*stack.Stack, *table.Table, types.VType) {
	tokens := p.Parse(code)
	return Run(tokens, stk, tbl)
}

// Temp.
func TableCall(name string, t *table.Table, s *stack.Stack) types.VType {
	var e types.VType = vnil.New()

	if n, ok := t.Native[name]; ok {
		_, _, e = Run(&n, s, t)

	} else if f, ok := t.Function[name]; ok {
		e = f(s, t)

	} else {
		a, _ := t.Aliases[name]
		return TableCall(a, t, s)
	}

	return e
}

func Run(tokens *p.Tokens, stk *stack.Stack, tbl *table.Table) (*stack.Stack, *table.Table, types.VType) {
	var val types.VType = vnil.New()

	for _, tok := range *tokens {
		switch tok.Key {
		case "str":
			stk.Push(vstring.New(tok.Val))

		case "int":
			stk.Push(vinteger.New(tok.Val))

		case "list":
			// TODO: need to make sure you can't add to ^tbl^ from a list.
			sub := stack.New()
			sub, _, _ = Eval(tok.Val, sub, tbl)
			stk.Push(vlist.New(sub))

		case "stm":
			stk.Push(vblock.New(tok.Val))

		case "fun":
			if tbl.Has(tok.Val) {
				val = TableCall(tok.Val, tbl, stk)
			} else {
				println("Unknown function: '" + tok.Val + "'")
			}

		default:
			println("Unknown token: " + tok.String())
		}
	}

	return stk, tbl, val
}
