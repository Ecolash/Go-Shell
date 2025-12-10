package repl

import (
	"fmt"
	"os"
	"path/filepath"
	"sort"
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

type builtinCompleter struct {
	lastPrefix string
	tabCount   int
}

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

	// reset tabCount when prefix changes
	if prefix != c.lastPrefix {
		c.tabCount = 0
	}
	c.lastPrefix = prefix

	// gather matching names
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
		fmt.Print("\a")
		return nil, 0
	}

	sort.Strings(matches)

	// if only one completion → autocomplete fully
	if len(matches) == 1 {
		c.tabCount = 0
		return [][]rune{[]rune(matches[0] + " ")}, len(prefix)
	}

	// multiple matches case
	c.tabCount++

	if c.tabCount == 1 {
		// first tab → ring bell
		fmt.Print("\a")
		return nil, len(prefix)
	}

	// second tab → print all
	c.tabCount = 0

	builder := strings.Join(matches, "  ") // two spaces

	// print results and restore prompt + buffer
	fmt.Printf("\n%s\n$ %s", builder, prefix)

	return nil, len(prefix)
}
