// Package parser implements a simple parser for the vodka language.
package parser

// Synopsis:
//
//   parser.Parse("1 2 + 'Hello World' [ '!' append ] 5 times")
//   ;=> [[fun: 1] [fun: 2] [fun: +] [str: 'Hello World'] [stm: '!' append]
//         [fun: 5] [fun: times]]

import (
	"regexp"
	"strings"
)

type Tokens []Token

// String returns a formatted string showing the tokens.
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

// Add appends a Token to the list of tokens.
func (ts *Tokens) Add(tok Token) {
	*ts = append(*ts, tok)
}

// AddToken appends a new Token with the key and value given.
func (ts *Tokens) AddToken(key, val string) {
	ts.Add(Token{key, val})
}

type Token struct {
	Key string
	Val string
}

func (t *Token) String() string {
	return "[" + t.Key + ": " + t.Val + "]"
}

// FullParse implements the parser for vodka source code. It takes a string of
// source code and returns the corresponding Tokens.
func FullParse(code string) *Tokens {
	list := new(Tokens)

	for i := 0; i < len(code); i++ {
		temp := ""

		switch c := code[i]; c {
		case '\t', ' ':
			// Ignore whitespace

		case '\n':
			list.AddToken("newline", "")

		case ';':
			i, temp = parseUntil(i, code, '\n')
			list.AddToken("comment",
				strings.TrimSpace(strings.TrimLeft(temp, ";")))

		case '.':
			list.AddToken("stm", "")

		case '\'':
			i++
			i, temp = parseUntil(i, code, '\'')
			list.AddToken("str", temp)

		case '"':
			i++
			i, temp = parseUntil(i, code, '"')
			list.AddToken("str", temp)

		case '(':
			i++
			i, temp = parseMatching(i, code, '(', ')')
			list.AddToken("list", temp)

		case '[':
			i++
			i, temp = parseMatching(i, code, '[', ']')
			list.AddToken("stm", temp)

		case '{':
			i++
			i, temp = parseMatching(i, code, '{', '}')
			list.AddToken("dict", temp)

		case ':':
			i++
			i, temp = parseUntilWhitespace(i, code)
			list.AddToken("stm", temp)

		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			i, temp = parseUntilWhitespace(i, code)
			matcher, _ := regexp.Compile("([0-9]+)\\.\\.(-?[0-9]+)")

			if matcher.MatchString(temp) {
				list.AddToken("range", temp)
			} else {
				list.AddToken("int", temp)
			}

		case '-':
			i, temp = parseUntilWhitespace(i, code)
			number, _ := regexp.Compile("-[0-9]+")
			ranger, _ := regexp.Compile("(-[0-9]+)\\.\\.(-?[0-9]+)")

			if ranger.MatchString(temp) {
				list.AddToken("range", temp)
			} else if number.MatchString(temp) {
				list.AddToken("int", temp)
			} else {
				list.AddToken("fun", temp)
			}

		default:
			i, temp = parseUntilWhitespace(i, code)
			list.AddToken("fun", temp)

		}
	}

	return list
}

// Parse implements a subset of FullParse, it will strip all comments and
// newlines from the returned Tokens.
func Parse(code string) *Tokens {
	list := new(Tokens)

	for _, tok := range *FullParse(code) {
		if tok.Key != "comment" && tok.Key != "newline" {
			list.Add(tok)
		}
	}

	return list
}

func parseUntil(idx int, code string, until uint8) (i int, s string) {
	return parseUntilAny(idx, code, []uint8{until})
}

func parseUntilWhitespace(idx int, code string) (i int, s string) {
	return parseUntilAny(idx, code, []uint8{' ', '\n', '\t'})
}

func parseUntilAny(idx int, code string, untils []uint8) (i int, s string) {
	str := ""

	for i := idx; i < len(code); i++ {
		c := code[i]
		for _, until := range untils {
			if c == until {
				return i, str
			}
		}
		str += string(c)
	}
	return len(code), str
}

func parseMatching(idx int, code string, op, cl uint8) (i int, s string) {
	str := ""

	for i := idx; i < len(code); i++ {
		c := code[i]
		if c == op {
			i++
			f := ""
			i, f = parseMatching(i, code, op, cl)
			str += string(op) + f + string(cl)
		} else if c == cl {
			return i, strings.TrimSpace(str)
		} else {
			str += string(c)
		}
	}

	return len(code), strings.TrimSpace(str)
}
