package repl

import (
	"fmt"
	"strings"

	"github.com/chzyer/readline"
	"github.com/codecrafters-io/shell-starter-go/internal/exec"
	"github.com/codecrafters-io/shell-starter-go/internal/parser"
)

func Start() {
	SetupHistory()
	rl, err := readline.NewEx(&readline.Config{
		Prompt:          "$ ",
		InterruptPrompt: "^C",
		EOFPrompt:       "exit",
		AutoComplete:    &builtinCompleter{},
		// HistoryFile:  "/tmp/my_shell_history.tmp",
	})
	if err != nil {
		fmt.Println("Failed to initialize readline:", err)
		return
	}
	defer rl.Close()
	for {
		input, err := rl.Readline()
		exec.ShellHistory = append(exec.ShellHistory, input)

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
		if command == exec.BuiltinExit {
			UpdateHistory()
		}
		exec.Run(command)
	}
}
