package main

import (
	"hawx.me/code/vodka/doc"
	"hawx.me/code/vodka/interpreter"
	"hawx.me/code/vodka/stack"
	"hawx.me/code/vodka/table"
	"hawx.me/code/vodka/types"
	"hawx.me/code/vodka/types/vnil"

	// "github.com/nsf/termbox-go"

	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

type Prompt struct {
	cursor string
	interp func(string) (string, string)
	r      *bufio.Reader
}

func NewPrompt(cursor string, interp func(string) (string, string)) *Prompt {
	return &Prompt{cursor, interp, bufio.NewReader(os.Stdin)}
}

func (p *Prompt) Loop() {
	for {
		line := p.promptLine()

		if line == "quit" {
			break
		}

		line = p.completeLine(line, "[", "]")
		line = p.completePair(line, "'")
		line = p.completePair(line, "\"")
		line = p.completeLine(line, "(", ")")

		stk, e := p.interp(line)
		fmt.Printf("%s => %s\n", stk, e)
	}
}

func (p *Prompt) promptLine() string {
	fmt.Print(p.cursor)
	l, _, _ := p.r.ReadLine()
	return p.toString(l)
}

func (p *Prompt) continueLine(depth int) string {
	s := ""
	for i := 0; i < depth; i++ {
		s += "  "
	}

	fmt.Print(".. " + s)
	l, _, _ := p.r.ReadLine()
	return p.toString(l)
}

func (p *Prompt) toString(bytes []uint8) string {
	str := ""
	for _, c := range bytes {
		str += string(byte(c))
	}
	return str
}

func (p *Prompt) completeLine(line, o, c string) string {
	opens := strings.Count(line, o)
	closes := strings.Count(line, c)

	for opens != closes {
		line += "\n" + p.continueLine(opens-closes)
		opens = strings.Count(line, o)
		closes = strings.Count(line, c)
	}

	return line
}

func (p *Prompt) completePair(line, o string) string {
	for strings.Count(line, o)%2 != 0 {
		line += "\n" + p.continueLine(0)
	}

	return line
}

var HELP_FLAGS []string = []string{"help", "-h", "--help", "-?"}

func isHelpFlag(s string) bool {
	for _, flag := range HELP_FLAGS {
		if flag == s {
			return true
		}
	}
	return false
}

func main() {
	stk := stack.New()
	tbl := interpreter.BootedTable(table.BOOT)

	if len(os.Args) > 1 {
		if os.Args[1] == "doc" {
			doc.Doc([]string{os.Args[2]}, "doc.html")

		} else if isHelpFlag(os.Args[1]) {
			fmt.Println(
				"Usage: vodka [files...]\n",
				"\n",
				"  Given no files to run, vodka will launch into a REPL.\n",
				"  Given a list of files, vodka will run each file in turn.\n",
			)

		} else {
			contents := ""
			for _, file := range os.Args[1:] {
				content, _ := ioutil.ReadFile(file)
				contents += string(content)
			}
			interpreter.Eval(contents, stk, tbl)
		}

	} else {
		fmt.Println("Vodka REPL, CTRL+C or type 'quit' to quit")

		interp := func(line string) (string, string) {
			var e types.VType = vnil.New()
			stk, tbl, e = interpreter.Eval(line, stk, tbl)
			return stk.TruncatedString(), e.String()
		}

		prompt := NewPrompt(">> ", interp)
		prompt.Loop()
	}
}
