package main

/*
 Synopsis:

 Parse("1 2 + 'Hello World' [ '!' append ] 5 times")
 //=> [[fun: 1] [fun: 2] [fun: +] [str: 'Hello World'] [stm: '!' append]
       [fun: 5] [fun: times]]
 */

import (
	"strings"
)

type Tokens []Token

func (ts *Tokens) String() string {
	str := "["
	for i, t := range *ts {
		if i > 0 {
			str += ", "
		}
		str += t.String()
	}
	return str + "]"
}

type Token struct {
	key  string
	val  string
}

func (t *Token) String() string {
	return "[" + t.key + ": " + t.val + "]"
}

func NewToken(key, val string) Token {
	return Token{key, val}
}



func Parse(code string) *Tokens {
	tokens := strings.Split(code, " ")
  list := new(Tokens)

	for i := 0; i < len(tokens); i++ {
		tok := tokens[i]

		if strings.HasPrefix(tok, "'") {
			idx, found := parseString(tokens, "'", i)
			i = idx
			*list = append(*list, NewToken("str", found))

		} else if strings.HasPrefix(tok, "\"") {
			idx, found := parseString(tokens, "\"", i)
			i = idx
			*list = append(*list, NewToken("str", found))

		} else if strings.TrimSpace(tok) ==  "[" {
			i++
			idx, found := parseBlock(tokens, i)
			i = idx
			*list = append(*list, NewToken("stm", found))

		} else if strings.HasPrefix(tok, "[") {
			idx, found := parseBlock(tokens, i)
			i = idx
			*list = append(*list, NewToken("stm", found))

		} else {
			*list = append(*list, NewToken("fun", tok))
		}
	}

	return list
}

/*
 Parses until a token ending in +until+ is found, returning what has been found
 as a string.
 */
func parseString(tokens []string, until string, idx int) (i int, str string) {
	str = ""
	for i := idx; i < len(tokens); i++ {
		tok := tokens[i]
		str += tok + " "
		if strings.HasSuffix(tok, until) {
			return i, str
		}
	}
	return len(tokens), strings.TrimSpace(str)
}

/*
 Parses a block, some tokens surrounded by square brackets.
 */
func parseBlock(tokens []string, idx int) (i int, str string) {
	str = ""
	for i := idx; i < len(tokens); i++ {
		tok := tokens[i]

		if strings.TrimSpace(tok) == "]" {
			return i, str
		}

		str += tok + " "
		if strings.HasSuffix(tok, "]") {
			return i, str
		}
	}
	return len(tokens), strings.TrimSpace(str)
}
