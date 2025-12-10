package repl

import (
	"fmt"
	"os"
	"path/filepath"
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

func loadPathExecutables() []string {
	path := os.Getenv("PATH")
	dirs := filepath.SplitList(path)
	seen := map[string]bool{}
	var out []string

	for _, d := range dirs {
		entries, err := os.ReadDir(d)
		if err != nil {
			continue
		}
		for _, e := range entries {
			if !e.Type().IsRegular() {
				continue
			}
			name := e.Name()
			if !seen[name] {
				seen[name] = true
				out = append(out, name)
			}
		}
	}
	return out
}

func (c *builtinCompleter) Do(line []rune, _ int) ([][]rune, int) {
	prefix := string(line)
	var matches []string

	for _, b := range builtinNames {
		if strings.HasPrefix(b, prefix) {
			matches = append(matches, b)
		}
	}

	for _, b := range loadPathExecutables() {
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
