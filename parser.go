package main


// Synopsis:
//
//   Parse("1 2 + 'Hello World' [ '!' append ] 5 times")
//   ;=> [[fun: 1] [fun: 2] [fun: +] [str: 'Hello World'] [stm: '!' append]
//         [fun: 5] [fun: times]]


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


func FullParse(code string) *Tokens {
	list := new(Tokens)

	for i := 0; i < len(code); i++ {
		temp := ""

		switch c := code[i]; c {
		case '\n', '\t', ' ':
			// Ignore whitespace

		case ';':
			i, temp = ParseUntil(i, code, '\n')
			*list = append(*list, NewToken("comment", strings.TrimSpace(strings.TrimLeft(temp, ";"))))

		case '.':
			*list = append(*list, NewToken("stm", ""))

		case '\'':
			i++
			i, temp = ParseUntil(i, code, '\'')
			*list = append(*list, NewToken("str", temp))

		case '"':
			i++
			i, temp = ParseUntil(i, code, '"')
			*list = append(*list, NewToken("str", temp))

		case '[':
			i++
			i, temp = ParseMatching(i, code, '[', ']')
			*list = append(*list, NewToken("stm", temp))

		case ':':
			i++
			i, temp = ParseUntilWhitespace(i, code)
			*list = append(*list, NewToken("stm", temp))

		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9', '-':
			i, temp = ParseUntilWhitespace(i, code)
			*list = append(*list, NewToken("int", temp))

		default:
			i, temp = ParseUntilWhitespace(i, code)
			*list = append(*list, NewToken("fun", temp))

		}
	}

	return list
}

func Parse(code string) *Tokens {
	list := new(Tokens)

	for _, tok := range *FullParse(code) {
		if tok.key != "comment" {
			*list = append(*list, tok)
		}
	}

	return list
}

func ParseUntil(idx int, code string, until uint8) (i int, s string) {
	return ParseUntilAny(idx, code, []uint8{until})
}

func ParseUntilWhitespace(idx int, code string) (i int, s string) {
	return ParseUntilAny(idx, code, []uint8{' ', '\n', '\t'})
}

func ParseUntilAny(idx int, code string, untils []uint8) (i int, s string) {
	str := ""

	for i := idx; i < len(code); i++ {
		c := code[i]
		for _, until := range untils {
			if c == until {
				return i, strings.TrimSpace(str)
			}
		}
		str += string(c)
	}
	return len(code), strings.TrimSpace(str)
}

func ParseMatching(idx int, code string, op, cl uint8) (i int, s string) {
	str := ""

	for i := idx; i < len(code); i++ {
		c := code[i]
		if c == op {
			i++
			f := ""
			i, f = ParseMatching(i, code, op, cl)
			str += "[" + f + "]"
		} else if c == cl {
			return i, strings.TrimSpace(str)
		} else {
			str += string(c)
		}
	}
	return len(code), strings.TrimSpace(str)
}
