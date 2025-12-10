package exec

import (
	"fmt"
	"os"

	"github.com/codecrafters-io/shell-starter-go/internal/parser"
)

func builtin(cmd *parser.Command) bool {
	switch cmd.Name {
	case "cd":
		if len(cmd.Args) == 0 {
			fmt.Println("cd: missing argument")
			return true
		}
		if err := os.Chdir(cmd.Args[0]); err != nil {
			fmt.Println("cd error:", err)
		}
		return true

	case "exit":
		os.Exit(0)
		return true
	}
	return false
}
