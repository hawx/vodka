package main

import (
	"io/ioutil"
)

// Generates documentation for the given list of input files, writing the output
// to the output directory given.
func Doc(input []string, output string) {
	println("Creating documentation...")

	for _, file := range input {
		println(parse(file).String())
	}
}

func parse(file string) *Tokens {
	contents, _ := ioutil.ReadFile(file)
	return FullParse(string(contents))
}
