package main

import (
	"strings"
	"bufio"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/hawx/vodka/interpreter"
	"github.com/hawx/vodka/stack"
	"github.com/hawx/vodka/types"
	"github.com/hawx/vodka/types/vnil"
	"github.com/hawx/vodka/doc"
)

func toString(bytes []uint8) string {
	str := ""
	for _, c := range bytes {
		str += string(byte(c))
	}
	return str
}

func promptLine(prompt string) string {
	fmt.Print(prompt)
	r := bufio.NewReader(os.Stdin)
	l, _, _ := r.ReadLine()
	return toString(l)
}

func completeLine(line, o, c string) string {
	opens  := strings.Count(line, o)
	closes := strings.Count(line, c)

	for opens != closes {
		s := ""
		for i := 0; i < (opens - closes); i++ {
			s += "  "
		}
		line += "\n" + promptLine(".. " + s)
		opens  = strings.Count(line, o)
		closes = strings.Count(line, c)
	}

	return line
}

func completePair(line, o string) string {
	for strings.Count(line, o) % 2 != 0 {
		line += "\n" + promptLine("..  ")
	}

	return line
}

func main() {
	stk := stack.New()
	tbl := interpreter.BootedTable()

	if len(os.Args) > 1 {
		if os.Args[1] == "doc" {
			doc.Doc([]string{os.Args[2]}, "doc.html")

		} else if os.Args[1] == "help" || os.Args[1] == "-h" || os.Args[1] == "--help" || os.Args[1] == "-?" {
			fmt.Println(
				"Usage: vodka [files...]\n",
				"\n",
				"  Given no files to run, vodka will launch into a REPL.\n",
				"  Given a list of files, vodka will run each file in turn.\n",
				)

		} else {
			for _, file := range os.Args[1:] {
				contents, _ := ioutil.ReadFile(file)
				interpreter.Eval(string(contents), stk, tbl)
			}
		}

	} else {
		fmt.Println("Vodka REPL, CTRL+C or type 'quit' to quit")

		for {
			line := promptLine(">> ")

			if line == "quit" {
				break
			}

			line = completeLine(line, "[", "]")
			line = completePair(line, "'")
			line = completePair(line, "\"")
			line = completeLine(line, "(", ")")

			var e types.VType = vnil.New()
			stk, tbl, e = interpreter.Eval(line, stk, tbl)
			fmt.Printf("%s => %s\n", stk.TruncatedString(), e.String())
		}
	}
}
