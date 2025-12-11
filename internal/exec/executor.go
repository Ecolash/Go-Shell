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
	pipes := make([][2]*os.File, n-1)
	for i := 0; i < n-1; i++ {
		r, w, err := os.Pipe()
		if err != nil {
			fmt.Println("pipe error:", err)
			return
		}
		pipes[i] = [2]*os.File{r, w}
	}

	getIO := func(i int) (stdin *os.File, stdout *os.File) {
		switch i {
		case 0:
			stdin = os.Stdin
		default:
			stdin = pipes[i-1][0]
		}
		switch i {
		case n - 1:
			stdout = os.Stdout
		default:
			stdout = pipes[i][1]
		}
		return
	}

	var process []*exec.Cmd
	for i, c := range p.Commands {
		stdin, stdout := getIO(i)
		// BUILTIN handling
		if Builtins[c.Name] {
			if c.Name == BuiltinCD || c.Name == BuiltinExit {
				fmt.Printf("%s: not allowed in pipeline\n", c.Name)
				return
			}
			origStdin := os.Stdin
			origStdout := os.Stdout
			os.Stdin = stdin
			os.Stdout = stdout
			builtin(c)

			os.Stdin = origStdin
			os.Stdout = origStdout
			continue
		}

		// EXTERNAL command
		cmd := exec.Command(c.Name, c.Args...)
		cmd.Stdin = stdin
		cmd.Stdout = stdout
		cmd.Stderr = os.Stderr

		if err := cmd.Start(); err != nil {
			fmt.Println(err)
			return
		}

		process = append(process, cmd)
	}
	for _, p := range pipes {
		_ = p[0].Close()
		_ = p[1].Close()
	}
	for _, cmd := range process {
		_ = cmd.Wait()
	}
}
