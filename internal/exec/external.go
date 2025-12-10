package exec

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/codecrafters-io/shell-starter-go/internal/parser"
)

func external(cmd *parser.Command, redirs []*Redirection) bool {
	c := exec.Command(cmd.Name, cmd.Args...)
	c.Stdin = os.Stdin
	c.Stdout = os.Stdout
	c.Stderr = os.Stderr
	applyRedirections(redirs, c)

	if err := c.Run(); err != nil {
		fmt.Println("error executing:", err)
	}
	return true
}
