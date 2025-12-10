package exec

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func doEcho(args []string) bool {
	fmt.Println(strings.Join(args, " "))
	return true
}

func doExit(args []string) bool {
	if len(args) == 0 {
		os.Exit(0)
		return true
	}
	return false
}

func doType(args []string) bool {
	if len(args) == 0 {
		fmt.Println("type: missing argument")
		return true
	}
	for _, word := range args {
		if Builtins[word] {
			fmt.Printf("%s is a shell builtin\n", word)
			continue
		}
		path, err := exec.LookPath(word)
		if err == nil {
			fmt.Printf("%s is %s\n", word, path)
			continue
		}
		fmt.Printf("%s: not found\n", word)
	}
	return true
}
