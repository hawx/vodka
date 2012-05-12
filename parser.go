package main

// Synopsis:
//
//   Parse("1 2 + 'Hello World' [ '!' append ] 5 times")
//   ;=> [[fun: 1] [fun: 2] [fun: +] [str: 'Hello World'] [stm: '!' append]
//         [fun: 5] [fun: times]]

import (
	"strings"
	"regexp"
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

func (ts *Tokens) Add(tok Token) {
	*ts = append(*ts, tok)
}

func (ts *Tokens) AddToken(key, val string) {
	ts.Add(NewToken(key, val))
}

type Token struct {
	key string
	val string
}

func (t *Token) String() string {
	return "[" + t.key + ": " + t.val + "]"
}

func NewToken(key, val string) Token {
	return Token{key, val}
}

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

		case ':':
			i++
			i, temp = parseUntilWhitespace(i, code)
			list.AddToken("stm", temp)

		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			i, temp = parseUntilWhitespace(i, code)
			list.AddToken("int", temp)

		case '-':
			i, temp = parseUntilWhitespace(i, code)
			matcher, _ := regexp.Compile("-[0-9]+")

			if matcher.MatchString(temp) {
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

func Parse(code string) *Tokens {
	list := new(Tokens)

	for _, tok := range *FullParse(code) {
		if tok.key != "comment" && tok.key != "newline" {
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
			str += "[" + f + "]"
		} else if c == cl {
			return i, strings.TrimSpace(str)
		} else {
			str += string(c)
		}
	}

	return len(code), strings.TrimSpace(str)
}
