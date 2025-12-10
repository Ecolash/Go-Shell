package parser

import "strings"

var doubleQuoteEscapes = map[byte]bool{
	'"':  true,
	'\\': true,
	'$':  true,
	'`':  true,
	'\n': true,
}

func Lex(input string) []string {
	var tokens []string
	var current strings.Builder
	inSingleQuote := false
	inDoubleQuote := false

	for i := 0; i < len(input); i++ {
		c := input[i]

		switch c {
		case '\\':
			if i+1 == len(input) {
				continue
			}
			if inSingleQuote {
				current.WriteByte(c)
				continue
			}
			if inDoubleQuote {
				next := input[i+1]
				if doubleQuoteEscapes[next] {
					current.WriteByte(next)
					i = i + 1
					continue
				}
				current.WriteByte(c)
				continue
			}
			c := input[i+1]
			current.WriteByte(c)
			i = i + 1
		case ' ':
			if inSingleQuote || inDoubleQuote {
				current.WriteByte(c)
			} else if current.Len() > 0 {
				tokens = append(tokens, current.String())
				current.Reset()
			}
		case '\'':
			if inDoubleQuote {
				current.WriteByte(c)
			} else {
				inSingleQuote = !inSingleQuote
			}
		case '"':
			if inSingleQuote {
				current.WriteByte(c)
			} else {
				inDoubleQuote = !inDoubleQuote
			}
		default:
			current.WriteByte(c)
		}
	}

	if current.Len() > 0 {
		tokens = append(tokens, current.String())
	}

	return tokens
}
