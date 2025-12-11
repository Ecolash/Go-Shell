package repl

import (
	"os"
	"strings"

	"github.com/codecrafters-io/shell-starter-go/internal/exec"
)

func SetupHistory() {
	histfile := os.Getenv("HISTFILE")
	if histfile == "" {
		histfile = "/tmp/my_shell_history.tmp"
	}
	exec.ReadHistory(histfile)
}

func WriteHistory() {
	histfile := os.Getenv("HISTFILE")
	if histfile == "" {
		histfile = "/tmp/my_shell_history.tmp"
	}
	exec.WriteHistory(histfile)
}

func AppendHistory() {
	histfile := os.Getenv("HISTFILE")
	if histfile == "" {
		histfile = "/tmp/my_shell_history.tmp"
	}
	exec.AppendHistory(histfile)
}

func UpdateHistory() {
	histfile := os.Getenv("HISTFILE")
	if histfile == "" {
		histfile = "/tmp/my_shell_history.tmp"
	}

	data, err := os.ReadFile(histfile)
	if err != nil {
		exec.WriteHistory(histfile)
		return
	}

	lines := strings.Split(string(data), "\n")
	if len(lines) == 0 || (len(lines) == 1 && lines[0] == "") {
		exec.WriteHistory(histfile)
		return
	}
	exec.AppendHistory(histfile)
	return
}
