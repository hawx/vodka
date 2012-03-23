package main

import (
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
	fmt.Println("Vodka REPL, CTRL+C or type 'quit' to quit")

	stk := NewStack()
	tbl := BootedTable()

	for {
		line := promptLine(">> ")

		if line == "quit" {
			break
		}

		stk, tbl = Eval(line, stk, tbl)
		fmt.Printf("=> %s\n", stk.String())
	}
}
