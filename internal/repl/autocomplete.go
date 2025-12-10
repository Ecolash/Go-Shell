package repl

import (
	"fmt"

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
	if len(completers) == 0 {
		fmt.Print("\x07") // beep to alert user
	}
	return completers
}
