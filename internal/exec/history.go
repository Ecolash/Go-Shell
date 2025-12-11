package exec

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var ShellHistory []string
var lastAppendIndex int

func printHistory(n int) {
	if n < 0 {
		n = len(ShellHistory)
	}
	total := len(ShellHistory)
	n = min(n, total)
	start := total - n
	start = max(0, start)
	for i := start; i < total; i++ {
		fmt.Printf("    %d  %s\n", i+1, ShellHistory[i])
	}
}

func ReadHistory(path string) {
	f, err := os.Open(path)
	if err != nil {
		// Bash behavior: history -r silently ignores missing file
		return
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}
		ShellHistory = append(ShellHistory, line)
	}
}

func WriteHistory(path string) {
	f, err := os.Create(path)
	if err != nil {
		fmt.Println("history: cannot write file:", path)
		return
	}
	defer f.Close()
	for _, line := range ShellHistory {
		_, _ = fmt.Fprintln(f, line)
	}
}

func AppendHistory(path string) {
	f, err := os.OpenFile(path, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("history: cannot append file:", path)
		return
	}
	defer f.Close()
	for i := lastAppendIndex; i < len(ShellHistory); i++ {
		_, _ = fmt.Fprintln(f, ShellHistory[i])
	}
	lastAppendIndex = len(ShellHistory)
}
