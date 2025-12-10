package repl

import (
	"fmt"
	"strings"

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

type builtinCompleter struct{}

func (c *builtinCompleter) Do(line []rune, _ int) (newLine [][]rune, length int) {
	prefix := string(line)
	var results [][]rune
	for _, b := range builtinNames {
		if strings.HasPrefix(b, prefix) {
			results = append(results, []rune(b))
		}
	}
	if len(results) == 0 {
		fmt.Print("\x07")
	}

	return results, len(prefix)
}
