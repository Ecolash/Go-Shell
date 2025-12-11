package exec

import "fmt"

var ShellHistory []string

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
