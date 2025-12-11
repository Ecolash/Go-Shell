package exec

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func doEcho(args []string) bool {
	args, redirs := parseRedirections(args)
	//fmt.Println(args)
	//fmt.Println(redirs[0].File)
	c := &exec.Cmd{
		Stdout: os.Stdout,
		Stderr: os.Stderr,
	}
	applyRedirections(redirs, c)
	_, err := fmt.Fprintln(c.Stdout, strings.Join(args, " "))
	if err != nil {
		_, _ = fmt.Fprintln(c.Stderr, "err")
		return false
	}
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

func doPwd(args []string) bool {
	if len(args) == 0 {
		cwd, err := os.Getwd()
		if err != nil {
			fmt.Println("pwd: error getting current directory:", err)
			return true
		}

		fmt.Println(cwd)
		return true
	}
	return false
}

func doCd(args []string) bool {
	if len(args) == 0 {
		fmt.Println("cd: missing argument")
		return true
	}
	dir := args[0]
	if args[0] == "~" { // ~ → home
		home, err := os.UserHomeDir()
		if err != nil {
			fmt.Println("cd: cannot determine home directory:", err)
			return true
		}
		dir = home
	}
	err := os.Chdir(dir)
	if err != nil {
		fmt.Println("cd:", strings.TrimSpace(args[0])+": No such file or directory")
	}
	return true
}

func doHistory(args []string) bool {
	if len(args) == 0 {
		printHistory()
		return true
	}
	return false
}
