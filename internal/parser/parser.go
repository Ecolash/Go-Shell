package parser

import "fmt"

func Parse(input string) (*Command, error) {
	tokens := Lex(input)
	if len(tokens) == 0 {
		return nil, fmt.Errorf("empty command")
	}

	return &Command{
		Name: tokens[0],
		Args: tokens[1:],
	}, nil
}
