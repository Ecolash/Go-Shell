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
	BuiltinType  = "type"
)

var Builtins = map[string]bool{
	BuiltinCD:    true,
	BuiltinExit:  true,
	BuiltinPWD:   true,
	BuiltinEcho:  true,
	BuiltinSet:   true,
	BuiltinUnset: true,
	BuiltinEnv:   true,
	BuiltinType:  true,
}

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

	case BuiltinType:
		args := cmd.Args
		if len(args) == 0 {
			fmt.Println("type: missing argument")
			return true
		}
		for _, word := range args {
			if Builtins[word] {
				fmt.Printf("%s is a shell builtin\n", word)
				continue
			}
			fmt.Printf("%s: not found\n", word)
		}
		return true
	}
	return false
}
