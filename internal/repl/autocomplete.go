package repl

import (
	"github.com/chzyer/readline"
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

func buildCompleters() []readline.PrefixCompleterInterface {
	var completers []readline.PrefixCompleterInterface
	for _, b := range builtinNames {
		completers = append(completers, readline.PcItem(b))
	}
	return completers
}
