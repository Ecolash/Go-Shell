package repl

import (
	"fmt"
	"os"
	"strings"

	"github.com/chzyer/readline"
	"github.com/codecrafters-io/shell-starter-go/internal/exec"
	"github.com/codecrafters-io/shell-starter-go/internal/parser"
)

func formatPrompt() string {
	cwd, err := os.Getwd()
	if err != nil {
		cwd = "?"
	}
	home := os.Getenv("HOME")
	if home != "" && strings.HasPrefix(cwd, home) {
		cwd = "~" + strings.TrimPrefix(cwd, home)
	}
	return fmt.Sprintf("%s $ ", cwd)
}

func Start() {
	histfile := os.Getenv("HISTFILE")
	if histfile == "" {
		histfile = "/tmp/my_shell_history.tmp"
	}

	SetupHistory()
	rl, err := readline.NewEx(&readline.Config{
		Prompt:          formatPrompt(),
		InterruptPrompt: "^C",
		EOFPrompt:       "exit",
		AutoComplete:    &builtinCompleter{},
		HistoryFile:     histfile,
	})
	if err != nil {
		fmt.Println("Failed to initialize readline:", err)
		return
	}
	defer rl.Close()

	for {
		rl.SetPrompt(formatPrompt())
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
