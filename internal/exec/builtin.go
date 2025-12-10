package exec

import (
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
		return doEcho(cmd.Args)
	case BuiltinExit:
		return doExit(cmd.Args)
	case BuiltinType:
		return doType(cmd.Args)
	}
	return false
}
