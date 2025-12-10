package repl

import (
	"fmt"
	"strings"

	"github.com/codecrafters-io/shell-starter-go/internal/exec"
	"github.com/codecrafters-io/shell-starter-go/internal/parser"
	"github.com/peterh/liner"
)

func Start() {
	line := liner.NewLiner()
	defer line.Close()
	line.SetCtrlCAborts(true)
	line.SetCompleter(func(input string) []string {
		var completions []string
		for _, b := range builtinNames {
			if strings.HasPrefix(b, input) {
				completions = append(completions, b)
			}
		}
		return completions
	})

	for {
		input, err := line.Prompt("$ ")
		if err != nil {
			fmt.Println("Error reading input:", err)
			return
		}

		input = strings.TrimSpace(input)
		if input == "" {
			continue
		}

		line.AppendHistory(input)
		command, err := parser.Parse(input)
		if err != nil {
			fmt.Println("parse error:", err)
			continue
		}
		exec.Run(command)
	}
}
