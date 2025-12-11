package repl

import (
	"os"

	"github.com/codecrafters-io/shell-starter-go/internal/exec"
)

func SetupHistory() {
	histfile := os.Getenv("HISTFILE")
	if histfile == "" {
		histfile = "/tmp/my_shell_history.tmp"
	}
	exec.ReadHistory(histfile)
}
