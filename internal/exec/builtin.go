package exec

import (
	"fmt"
	"os"
	"strings"

	"github.com/codecrafters-io/shell-starter-go/internal/parser"
)

const (
	BuiltinCD    = "cd"
	BuiltinExit  = "exit"
	BuiltinPWD   = "pwd"
	BuiltinEcho  = "echo"
	BuiltinSet   = "set"
	BuiltinUnset = "unset"
	BuiltinEnv   = "env"
)

func builtin(cmd *parser.Command) bool {
	switch cmd.Name {
	//case BuiltinCD:
	//	if len(cmd.Args) == 0 {
	//		fmt.Println("cd: missing argument")
	//		return true
	//	}
	//	if err := os.Chdir(cmd.Args[0]); err != nil {
	//		fmt.Println("cd error:", err)
	//	}
	//	return true

	case BuiltinEcho:
		fmt.Println(strings.Join(cmd.Args, " "))
		return true

	case BuiltinExit:
		os.Exit(0)
		return true
	}
	return false
}
