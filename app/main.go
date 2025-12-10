package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Ensures gofmt doesn't remove the "fmt" import in stage 1 (feel free to remove this!)
var _ = fmt.Print

func main() {
	fmt.Print("$ ")
	reader := bufio.NewReader(os.Stdin)
	command, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println(err)
	}

	command = strings.TrimSuffix(command, "\n")
	fmt.Println(command + ": command not found")
}
