package interpreter

import (
	"fmt"
	p "github.com/hawx/vodka/parser"
	"io/ioutil"

	"github.com/hawx/vodka/stack"
	"github.com/hawx/vodka/table"

	"github.com/hawx/vodka/types"
	"github.com/hawx/vodka/types/vblock"
	"github.com/hawx/vodka/types/vboolean"
	"github.com/hawx/vodka/types/vinteger"
	"github.com/hawx/vodka/types/vlist"
	"github.com/hawx/vodka/types/vnil"
	"github.com/hawx/vodka/types/vrange"
	"github.com/hawx/vodka/types/vstring"
)

// Special assign that will be executed before the interpreter exits.
var onExitStms *p.Tokens

// BootedTable returns a table with built in functions defined.
func BootedTable(boot string) *table.Table {
	tbl := table.New()

	tbl.Define("eval", func(s *stack.Stack, t *table.Table) types.VType {
		str := s.Pop().Value().(string)
		_, _, v := eval(str, s, t)
		return v
	})

	tbl.Define("alias", func(s *stack.Stack, t *table.Table) types.VType {
		from := s.Pop().Value().(string)
		to := s.Pop().Value().(string)
		t.Alias(from, to)
		return vnil.New()
	})

	tbl.Define("define", func(s *stack.Stack, t *table.Table) types.VType {
		stms := s.Pop().Value().(*p.Tokens)
		name := s.Pop().Value().(string)
		t.DefineNative(name, stms)
		return vnil.New()
	})

	tbl.Define("on-exit", func(s *stack.Stack, t *table.Table) types.VType {
		onExitStms = s.Pop().Value().(*p.Tokens)
		return vnil.New()
	})

	tbl.Define("defined", func(s *stack.Stack, t *table.Table) types.VType {
		defined := t.Defined()
		list := make([]types.VType, len(defined))
		for i, name := range defined {
			list[i] = vstring.New(name)
		}

		s.Push(vlist.NewFromList(list))
		return vnil.New()
	})

	tbl.Define("type", func(s *stack.Stack, t *table.Table) types.VType {
		v := vstring.New(s.Pop().Type())
		s.Push(v)
		return vnil.New()
	})

	// Types

	tbl.Define("integer", func(s *stack.Stack, t *table.Table) types.VType {
		v := s.Pop().Value().(string)
		s.Push(vinteger.New(v))
		return vnil.New()
	})

	tbl.Define("string", func(s *stack.Stack, t *table.Table) types.VType {
		v := s.Pop()

		if v.Type() == "string" {
			s.Push(v)
		} else {
			s.Push(vstring.New(v.String()))
		}

		return vnil.New()
	})

	tbl.Define("list", func(s *stack.Stack, t *table.Table) types.VType {
		v := s.Pop()

		if r,ok := v.(*vrange.VRange); ok {
			list := r.List()
			s.Push(list)

		} else {
			list := make([]types.VType, 1)
			list[0] = v

			s.Push(vlist.NewFromList(list))
		}

		return vnil.New()
	})

	tbl.Define("range", func(s *stack.Stack, t *table.Table) types.VType {
		start := s.Pop().(types.Rangeable)
		end := s.Pop().(types.Rangeable)

		s.Push(vrange.NewFromStartAndEnd(start, end))

		return vnil.New()
	})

	tbl.Define("max", func(s *stack.Stack, t *table.Table) types.VType {
		v := s.Pop().(*vrange.VRange)

		s.Push(v.Max())

		return vnil.New()
	})

	tbl.Define("min", func(s *stack.Stack, t *table.Table) types.VType {
		v := s.Pop().(*vrange.VRange)

		s.Push(v.Min())

		return vnil.New()
	})

	// I/O

	tbl.Define("print", func(s *stack.Stack, t *table.Table) types.VType {
		v := s.Pop().Value().(string)
		fmt.Println(v)
		return vnil.New()
	})

	tbl.Define("p", func(s *stack.Stack, t *table.Table) types.VType {
		v := s.PopString().String()
		fmt.Println(v[1 : len(v)-1])
		return vnil.New()
	})

	tbl.Define("read", func(s *stack.Stack, t *table.Table) types.VType {
		contents, _ := ioutil.ReadFile(s.Pop().Value().(string))
		str := vstring.New(string(contents))
		s.Push(str)
		return vnil.New()
	})

	// Stack operations

	tbl.Define("pop", func(s *stack.Stack, t *table.Table) types.VType {
		v := s.Pop()
		return v
	})

	tbl.Define("size", func(s *stack.Stack, t *table.Table) types.VType {
		v := vinteger.NewFromInt(s.Size())
		s.Push(v)
		return vnil.New()
	})

	tbl.Define("dup", func(s *stack.Stack, t *table.Table) types.VType {
		v := s.Top()
		s.Push(v)
		return vnil.New()
	})

	tbl.Define("swap", func(s *stack.Stack, t *table.Table) types.VType {
		a := s.Pop()
		b := s.Pop()
		s.Push(a)
		s.Push(b)
		return vnil.New()
	})

	tbl.Define("drop", func(s *stack.Stack, t *table.Table) types.VType {
		s.Clear()
		return vnil.New()
	})

	tbl.Define("compose", func(s *stack.Stack, t *table.Table) types.VType {
		a := s.Pop()
		b := s.Pop()
		c := vblock.New(b.(*vblock.VBlock).BareValue() + " " + a.(*vblock.VBlock).BareValue())
		s.Push(c)
		return vnil.New()
	})

	tbl.Define("wrap", func(s *stack.Stack, t *table.Table) types.VType {
		b := s.Pop()
		r := vblock.New(b.String())
		s.Push(r)
		return vnil.New()
	})

	// Arithmetic

	tbl.Define("add", func(s *stack.Stack, t *table.Table) types.VType {
		add := vinteger.NewFromInt(s.Pop().Value().(int) + s.Pop().Value().(int))
		s.Push(add)
		return vnil.New()
	})

	tbl.Define("mult", func(s *stack.Stack, t *table.Table) types.VType {
		mult := vinteger.NewFromInt(s.Pop().Value().(int) * s.Pop().Value().(int))
		s.Push(mult)
		return vnil.New()
	})

	tbl.Define("sub", func(s *stack.Stack, t *table.Table) types.VType {
		a := s.Pop().Value().(int)
		b := s.Pop().Value().(int)
		sub := vinteger.NewFromInt(a - b)
		s.Push(sub)
		return vnil.New()
	})

	tbl.Define("div", func(s *stack.Stack, t *table.Table) types.VType {
		a := s.Pop().Value().(int)
		b := s.Pop().Value().(int)
		div := vinteger.NewFromInt(a / b)
		s.Push(div)
		return vnil.New()
	})

	tbl.Define("neg", func(s *stack.Stack, t *table.Table) types.VType {
		val := vinteger.NewFromInt(-s.Pop().Value().(int))
		s.Push(val)
		return vnil.New()
	})

	// Logical

	tbl.Define("true", func(s *stack.Stack, t *table.Table) types.VType {
		s.Push(vboolean.True())
		return vnil.New()
	})

	tbl.Define("false", func(s *stack.Stack, t *table.Table) types.VType {
		s.Push(vboolean.False())
		return vnil.New()
	})

	tbl.Define("nil", func(s *stack.Stack, t *table.Table) types.VType {
		s.Push(vnil.New())
		return vnil.New()
	})

	tbl.Define("or", func(s *stack.Stack, t *table.Table) types.VType {
		a := s.Pop().Value().(bool)
		b := s.Pop().Value().(bool)
		val := vboolean.False()
		if a || b {
			val = vboolean.True()
		}
		s.Push(val)
		return vnil.New()
	})

	tbl.Define("and", func(s *stack.Stack, t *table.Table) types.VType {
		a := s.Pop().Value().(bool)
		b := s.Pop().Value().(bool)
		val := vboolean.False()
		if a && b {
			val = vboolean.True()
		}
		s.Push(val)
		return vnil.New()
	})

	tbl.Define("compare", func(s *stack.Stack, t *table.Table) types.VType {
		a := s.Pop()
		b := s.Pop()
		val := vinteger.NewFromInt(a.Compare(b))
		s.Push(val)
		return vnil.New()
	})

	tbl.Define("eq?", func(s *stack.Stack, t *table.Table) types.VType {
		a := s.Pop()
		b := s.Pop()
		val := vboolean.New(a.Compare(b) == 0)
		s.Push(val)
		return vnil.New()
	})

	// Flow

	tbl.Define("if-else", func(s *stack.Stack, t *table.Table) types.VType {
		a := s.Pop().Value().(*p.Tokens)
		b := s.Pop().Value().(*p.Tokens)
		cond := s.Pop().Value().(bool)
		if cond {
			s, t, _ = run(a, s, t)
		} else {
			s, t, _ = run(b, s, t)
		}
		return vnil.New()
	})

	tbl.Define("call", func(s *stack.Stack, t *table.Table) types.VType {
		val := s.Top().Value()
		switch val.(type) {
		case *p.Tokens:
			s.Pop()
			run(val.(*p.Tokens), s, t)
		case *vblock.VBlock:
			toks := new(p.Tokens)
			*toks = append(*toks, p.Token{"fun", "call"})
			run(toks, s, t)
		default:
			println("Unexpected type")
		}
		return vnil.New()
	})

	tbl.Define("without", func(s *stack.Stack, t *table.Table) types.VType {
		save := s.Pop()
		tokens := new(p.Tokens)
		*tokens = append(*tokens, p.Token{"fun", "call"})
		run(tokens, s, t)
		s.Push(save)
		return vnil.New()
	})

	tbl.Define("without2", func(s *stack.Stack, t *table.Table) types.VType {
		save1 := s.Pop()
		save2 := s.Pop()
		tokens := new(p.Tokens)
		*tokens = append(*tokens, p.Token{"fun", "call"})
		run(tokens, s, t)
		s.Push(save2)
		s.Push(save1)
		return vnil.New()
	})

	// Strings

	tbl.Define("concat", func(s *stack.Stack, t *table.Table) types.VType {
		a := s.Pop().Value().(string)
		b := s.Pop().Value().(string)
		c := vstring.New(b + a)
		s.Push(c)
		return vnil.New()
	})

	// Lists

	tbl.Define("head", func(s *stack.Stack, t *table.Table) types.VType {
		h := s.Pop().Value().([]types.VType)
		if len(h) > 0 {
			s.Push(h[0])
		} else {
			s.Push(vnil.New())
		}
		return vnil.New()
	})

	tbl.Define("tail", func(s *stack.Stack, t *table.Table) types.VType {
		v := s.Pop().Value().([]types.VType)
		if len(v) > 0 {
			s.Push(vlist.NewFromList(v[1:]))
		} else {
			s.Push(vnil.New())
		}
		return vnil.New()
	})

	tbl.Define("cons", func(s *stack.Stack, t *table.Table) types.VType {
		v := s.Pop().(types.VType)
		l := s.Pop().Value().([]types.VType)
		l = append(l, v)
		s.Push(vlist.NewFromList(l))
		return vnil.New()
	})

	tbl.Define("append", func(s *stack.Stack, t *table.Table) types.VType {
		a := s.Pop().Value().([]types.VType)
		b := s.Pop().Value().([]types.VType)

		// Ripped from http://golang.org/doc/effective_go.html#slices
		l := len(a)
		if l+len(b) > cap(a) {
			// Allocate double what's needed, for future growth.
			newSlice := make([]types.VType, (l+len(b))*2)
			copy(newSlice, a)
			a = newSlice
		}
		a = a[0 : l+len(b)]
		for i, c := range b {
			a[l+i] = c
		}

		s.Push(vlist.NewFromList(a))
		return vnil.New()
	})

	tbl.Define("apply", func(s *stack.Stack, t *table.Table) types.VType {
		f := s.Pop().Value().(*p.Tokens)
		l := s.Pop().Value().([]types.VType)

		stk := make(stack.Stack, len(l))
		for i, o := range l {
			stk[i] = o.(types.VType)
		}

		newstk, _, v := run(f, &stk, t)

		list := make([]types.VType, len(*newstk))
		for i, o := range *newstk {
			list[i] = o.(types.VType)
		}
		s.Push(vlist.NewFromList(list))

		return v
	})

	tbl.Define("reverse", func(s *stack.Stack, t *table.Table) types.VType {
		l := s.Pop().Value().([]types.VType)

		for i, j := 0, len(l)-1; i < j; i, j = i+1, j-1 {
			l[i], l[j] = l[j], l[i]
		}

		s.Push(vlist.NewFromList(l))

		return vnil.New()
	})

	// spec

	var specTbl map[string]bool
	passCount := 0
	failCount := 0

	tbl.Define("describe", func(s *stack.Stack, t *table.Table) types.VType {
		block := s.Pop().Value().(*p.Tokens)
		desc := s.Pop().Value().(string)

		fmt.Print("\n" + desc + "\n  ")

		specTbl = map[string]bool{}
		run(block, s, t)

		fmt.Print("\n")

		for v, k := range specTbl {
			if k {
				passCount++
			} else {
				failCount++
				fmt.Println("  FAIL " + v)
			}
		}

		onExitStms = p.Parse("'" + fmt.Sprintf("\n  %v pass / %v fail", passCount, failCount) + "' print")

		return vnil.New()
	})

	tbl.Define("can", func(s *stack.Stack, t *table.Table) types.VType {
		block := s.Pop().Value().(*p.Tokens)
		desc := s.Pop().Value().(string)

		// IMPORTANT: run on an empty stack each time
		ns, _, _ := run(block, stack.New(), t)
		if ns.Pop().Value().(bool) {
			fmt.Print(".")
			specTbl[desc] = true
		} else {
			fmt.Print("x")
			specTbl[desc] = false
		}

		return vnil.New()
	})

	_, tbl, _ = eval(boot, stack.New(), tbl)

	return tbl
}

func eval(code string, stk *stack.Stack, tbl *table.Table) (*stack.Stack, *table.Table, types.VType) {
	tokens := p.Parse(code)
	return run(tokens, stk, tbl)
}

// tableCall finds the named function in the Table and runs it using the Stack
// given. It returns the value returned by the function.
func tableCall(name string, t *table.Table, s *stack.Stack) types.VType {
	var e types.VType = vnil.New()

	if n, ok := t.Native[name]; ok {
		_, _, e = run(&n, s, t)

	} else if f, ok := t.Function[name]; ok {
		e = f.Apply(s, t)

	} else {
		a, _ := t.Aliases[name]
		return tableCall(a, t, s)
	}

	return e
}

func run(tokens *p.Tokens, stk *stack.Stack, tbl *table.Table) (*stack.Stack, *table.Table, types.VType) {
	var val types.VType = vnil.New()

	for _, tok := range *tokens {
		switch tok.Key {
		case "str": // strings are pushed onto the stack as vstrings
			stk.Push(vstring.New(tok.Val))

		case "int": // ints are pushed onto the stack as vintegers
			stk.Push(vinteger.New(tok.Val))

		case "range": // ranges are pushed onto the stack as vranges
			stk.Push(vrange.New(tok.Val))

		case "list": // lists are pushed onto the stack as vlists
			sub := stack.New()
			sub, _, _ = eval(tok.Val, sub, tbl)
			stk.Push(vlist.New(sub))

		case "stm": // statements are pushed onto the stack as vblocks
			stk.Push(vblock.New(tok.Val))

		case "fun": // functions are called immediately
			if tbl.Has(tok.Val) {
				val = tableCall(tok.Val, tbl, stk)
			} else {
				println("Unknown function: '" + tok.Val + "'")
			}

		default:
			println("Unknown token: " + tok.String())
		}
	}

	return stk, tbl, val
}

// Eval runs the string of vodka code given using the Stack and Table passed. It
// returns the modified Stack and Table, along with the last returned value.
func Eval(code string, stk *stack.Stack, tbl *table.Table) (*stack.Stack, *table.Table, types.VType) {
	s, t, v := eval(code, stk, tbl)

	if onExitStms != nil {
		run(onExitStms, s, t)
	}

	return s, t, v
}

// Run takes a series of tokens and runs them using the Stack and Table
// given. It returns the modified Stack and Table along with the last returned
// value.
func Run(tokens *p.Tokens, stk *stack.Stack, tbl *table.Table) (*stack.Stack, *table.Table, types.VType) {
	s, t, v := run(tokens, stk, tbl)

	if onExitStms != nil {
		run(onExitStms, s, t)
	}

	return s, t, v
}
