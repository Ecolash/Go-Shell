package parser

import (
	"fmt"
)

func Parse(input string) (interface{}, error) {
	tokens := Lex(input)
	if len(tokens) == 0 {
		return nil, fmt.Errorf("empty command")
	}

	// Detect pipes
	var segments [][]string
	var current []string

	for _, tok := range tokens {
		if tok == "|" {
			if len(current) == 0 {
				return nil, fmt.Errorf("syntax error near |")
			}
			segments = append(segments, current)
			current = []string{}
		} else {
			current = append(current, tok)
		}
	}
	segments = append(segments, current)

	// Single command
	if len(segments) == 1 {
		seg := segments[0]
		return &Command{
			Name: seg[0],
			Args: seg[1:],
		}, nil
	}

	// Multi-command pipeline
	p := &Pipeline{}
	for _, seg := range segments {
		cmd := &Command{
			Name: seg[0],
			Args: seg[1:],
		}
		p.Commands = append(p.Commands, cmd)
	}
	return p, nil
}
