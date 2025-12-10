package parser

func Lex(input string) []string {
	var tokens []string
	current := ""

	for _, c := range input {
		if c == ' ' || c == '\n' || c == '\t' {
			if current != "" {
				tokens = append(tokens, current)
				current = ""
			}
			continue
		}
		current += string(c)
	}

	if current != "" {
		tokens = append(tokens, current)
	}
	return tokens
}
