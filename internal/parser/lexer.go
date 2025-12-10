package parser

import "strings"

func Lex(input string) []string {
	var tokens []string
	var current strings.Builder
	inSingleQuote := false
	inDoubleQuote := false

	for i := 0; i < len(input); i++ {
		c := input[i]

		switch c {
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
