package repl

import (
	"fmt"
	"strings"

	"github.com/chzyer/readline"
	"github.com/codecrafters-io/shell-starter-go/internal/exec"
	"github.com/codecrafters-io/shell-starter-go/internal/parser"
)

func Start() {
	rl, err := readline.NewEx(&readline.Config{
		Prompt:          "$ ",
		InterruptPrompt: "^C",
		EOFPrompt:       "exit",
		HistoryFile:     "/tmp/my_shell_history.tmp",
		AutoComplete:    &builtinCompleter{},
	})
	if err != nil {
		fmt.Println("Failed to initialize readline:", err)
		return
	}
	defer rl.Close()
	for {
		input, err := rl.Readline()
		if err != nil { // EOF or Ctrl+C
			break
		}
		input = strings.TrimSpace(input)
		if input == "" {
			continue
		}

		command, err := parser.Parse(input)
		if err != nil {
			fmt.Println("parse error:", err)
			continue
		}
		exec.Run(command)
	}
}
