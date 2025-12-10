package main

import (
	"fmt"

	"github.com/codecrafters-io/shell-starter-go/internal/repl"
)

// Ensures gofmt doesn't remove the "fmt" import in stage 1 (feel free to remove this!)
var _ = fmt.Print

func main() {
	repl.Start()
}
