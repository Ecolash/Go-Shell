package repl

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/codecrafters-io/shell-starter-go/internal/exec"
	"github.com/codecrafters-io/shell-starter-go/internal/parser"
)

func Start() {
	for {
		fmt.Print("$ ")
		reader := bufio.NewReader(os.Stdin)
		line, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading input:", err)
			os.Exit(1)
		}

		line = strings.TrimSuffix(line, "\n")
		command, err := parser.Parse(line)
		if err != nil {
			fmt.Println("parse error:", err)
			continue
		}
		exec.Run(command)
	}
}
