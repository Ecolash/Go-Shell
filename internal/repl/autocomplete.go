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

func (c *builtinCompleter) Do(line []rune, _ int) ([][]rune, int) {
	prefix := string(line)
	var matches []string

	for _, b := range builtinNames {
		if strings.HasPrefix(b, prefix) {
			matches = append(matches, b)
		}
	}

	if len(matches) == 0 {
		fmt.Print("\x07")
		return nil, 0
	}

	if len(matches) == 1 {
		match := matches[0]
		missing := match[len(prefix):] + " "
		return [][]rune{[]rune(missing)}, len(prefix)
	}

	results := make([][]rune, len(matches))
	for i, m := range matches {
		results[i] = []rune(m)
	}
	return results, len(prefix)
}
