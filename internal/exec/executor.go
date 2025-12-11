package exec

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/codecrafters-io/shell-starter-go/internal/parser"
)

func Run(node interface{}) {
	switch n := node.(type) {
	case *parser.Command:
		runSingleCommand(n)
	case *parser.Pipeline:
		runPipeline(n)
	default:
		fmt.Println("unknown parse type")
	}
}

func runSingleCommand(cmd *parser.Command) {
	if builtin(cmd) {
		return
	}
	_, err := exec.LookPath(cmd.Name)
	args, redirs := parseRedirections(cmd.Args)
	cmd.Args = args
	if err != nil {
		fmt.Println(cmd.Name + ": command not found")
		return
	}
	external(cmd, redirs)
}

func runPipeline(p *parser.Pipeline) {
	n := len(p.Commands)
	procs := make([]*exec.Cmd, n)
	pipes := make([][2]*os.File, n-1)
	for i, c := range p.Commands {
		procs[i] = exec.Command(c.Name, c.Args...)
		procs[i].Stderr = os.Stderr
	}

	for i := 0; i < n-1; i++ {
		r, w, _ := os.Pipe()
		pipes[i] = [2]*os.File{r, w}
		procs[i].Stdout = w  // left -> write
		procs[i+1].Stdin = r // right <- read
	}

	procs[0].Stdin = os.Stdin
	procs[n-1].Stdout = os.Stdout
	for i := 0; i < n; i++ {
		_ = procs[i].Start()
	}

	// Parent closes pipes
	for i := 0; i < n-1; i++ {
		_ = pipes[i][0].Close()
		_ = pipes[i][1].Close()
	}

	_ = procs[n-1].Wait()
	for i := 0; i < n-1; i++ {
		_ = procs[i].Wait()
	}
}
