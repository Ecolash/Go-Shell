package exec

import (
	"fmt"
	"os/exec"

	"github.com/codecrafters-io/shell-starter-go/internal/parser"
)

func Run(cmd *parser.Command) {
	if builtin(cmd) {
		return
	}
	_, err := exec.LookPath(cmd.Name)
	if err != nil {
		fmt.Println(cmd.Name + ": command not found")
		return
	}
	external(cmd)
}
