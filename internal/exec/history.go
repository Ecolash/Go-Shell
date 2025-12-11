package exec

import "fmt"

var ShellHistory []string

func printHistory() {
	for i, entry := range ShellHistory {
		fmt.Printf("    %d  %s\n", i+1, entry)
	}
}
