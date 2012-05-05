package main

import (
	"strings"
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
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
	if len(os.Args) > 1 {
		if os.Args[1] == "doc" {
			Doc([]string{os.Args[2]}, "doc.html")

		} else {
			for _, file := range os.Args[1:] {
				contents, _ := ioutil.ReadFile(file)
				BareEval(string(contents))
			}
		}

	} else {
		fmt.Println("Vodka REPL, CTRL+C or type 'quit' to quit")
		stk := NewStack()
		tbl := BootedTable()

		for {
			line := promptLine(">> ")

			if line == "quit" {
				break
			}

			line = completeLine(line, "[", "]")
			line = completePair(line, "'")
			line = completePair(line, "\"")

			var e VType = VNil()
			stk, tbl, e = Eval(line, stk, tbl)
			fmt.Printf("%s => %s\n", stk.TruncatedString(), e.String())
		}
	}
}
