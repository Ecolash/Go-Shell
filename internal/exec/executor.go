package exec

import (
	"fmt"

	"github.com/codecrafters-io/shell-starter-go/internal/parser"
)

func Run(cmd *parser.Command) {
	if builtin(cmd) {
		return
	}

	fmt.Println(cmd.Name + ": command not found")

	//c := exec.Command(cmd.Name, cmd.Args...)
	//c.Stdin = os.Stdin
	//c.Stdout = os.Stdout
	//c.Stderr = os.Stderr
	//
	//if err := c.Run(); err != nil {
	//	fmt.Println("error:", err)
	//}
}
