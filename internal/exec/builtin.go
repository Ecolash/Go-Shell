package exec

import (
	"github.com/codecrafters-io/shell-starter-go/internal/parser"
)

const (
	BuiltinCD      = "cd"
	BuiltinExit    = "exit"
	BuiltinPWD     = "pwd"
	BuiltinEcho    = "echo"
	BuiltinHistory = "history"
	BuiltinSet     = "set"
	BuiltinUnset   = "unset"
	BuiltinEnv     = "env"
	BuiltinType    = "type"
)

var Builtins = map[string]bool{
	BuiltinCD:      true,
	BuiltinExit:    true,
	BuiltinPWD:     true,
	BuiltinEcho:    true,
	BuiltinHistory: true,
	BuiltinSet:     true,
	BuiltinUnset:   true,
	BuiltinEnv:     true,
	BuiltinType:    true,
}

func builtin(cmd *parser.Command) bool {
	switch cmd.Name {
	case BuiltinCD:
		return doCd(cmd.Args)
	case BuiltinEcho:
		return doEcho(cmd.Args)
	case BuiltinExit:
		return doExit(cmd.Args)
	case BuiltinType:
		return doType(cmd.Args)
	case BuiltinHistory:
		return doHistory(cmd.Args)
	case BuiltinPWD:
		return doPwd(cmd.Args)
	}
	return false
}
