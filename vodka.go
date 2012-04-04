package main

import (
	"io/ioutil"
	"fmt"
	"os"
	"bufio"
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

func main() {
	if len(os.Args) == 2 {
		if os.Args[1] == "doc" {
			Doc([]string{"boot.vk"}, "doc")
		} else {
			contents, _ := ioutil.ReadFile(os.Args[1])
			BareEval(string(contents))
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

			var e VType = VNil()
			stk, tbl, e = Eval(line, stk, tbl)
			fmt.Printf("%s => %s\n", stk.TruncatedString(), e.String())
		}
	}
}
