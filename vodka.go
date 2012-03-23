package main

import (
	"fmt"
)

func main() {
	code := "1 2 3 + eq? 'Hello World' \"Hello Josh\" [ dup inc ] 5 times"
	fmt.Println(Parse(code))

	BareEval("1 2 3 + *")
}
