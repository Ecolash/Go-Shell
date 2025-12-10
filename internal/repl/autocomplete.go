package repl

import (
	"github.com/codecrafters-io/shell-starter-go/internal/exec"
)

var builtinNames = []string{
	exec.BuiltinCD,
	exec.BuiltinExit,
	exec.BuiltinPWD,
	exec.BuiltinEcho,
	exec.BuiltinSet,
	exec.BuiltinUnset,
	exec.BuiltinEnv,
	exec.BuiltinType,
}
