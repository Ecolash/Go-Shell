package exec

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
)

type Redirection struct {
	Fd     int    // file descriptor: 1 = stdout, 2 = stderr
	File   string // target filename
	Append bool   // true for >>
}

func parseRedirections(args []string) ([]string, []*Redirection) {
	var newArgs []string
	var redirs []*Redirection

	i := 0
	for i < len(args) {
		arg := args[i]
		var fd int
		var isAppend bool
		var file string
		matched := false

		if strings.HasPrefix(arg, ">>") {
			fd = 1
			isAppend = true
			if len(arg) > 2 {
				file = arg[2:]
			} else if i+1 < len(args) {
				file = args[i+1]
				i++
			} else {
				fmt.Println("syntax error near unexpected token `newline`")
				return nil, nil
			}
			matched = true
		} else if strings.HasPrefix(arg, ">") {
			fd = 1
			isAppend = false
			if len(arg) > 1 {
				file = arg[1:]
			} else if i+1 < len(args) {
				file = args[i+1]
				i++
			} else {
				fmt.Println("syntax error near unexpected token `newline`")
				return nil, nil
			}
			matched = true
		} else if strings.Contains(arg, ">>") {
			parts := strings.SplitN(arg, ">>", 2)
			if len(parts) == 2 {
				fd, _ = strconv.Atoi(parts[0])
				file = parts[1]
				isAppend = true
				matched = true
			}
		} else if strings.Contains(arg, ">") {
			parts := strings.SplitN(arg, ">", 2)
			if len(parts) == 2 {
				fd, _ = strconv.Atoi(parts[0])
				file = parts[1]
				isAppend = false
				matched = true
			}
		}

		if matched {
			redirs = append(redirs, &Redirection{Fd: fd, File: file, Append: isAppend})
		} else {
			newArgs = append(newArgs, arg)
		}
		i++
	}

	return newArgs, redirs
}

func applyRedirections(redirs []*Redirection, cmd *exec.Cmd) {
	for _, r := range redirs {
		var f *os.File
		var err error
		dir := filepath.Dir(r.File)
		err = os.MkdirAll(dir, 0755)
		if err != nil {
			fmt.Println("mkdir err:", err)
			return
		}
		if r.Append {
			f, err = os.OpenFile(r.File, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
		} else {
			f, err = os.Create(r.File)
		}
		if err != nil {
			fmt.Println("cannot open file:", r.File)
			continue
		}

		switch r.Fd {
		case 1:
			cmd.Stdout = f
		case 2:
			cmd.Stderr = f
		default:
			fmt.Println("unsupported fd redirection:", r.Fd)
		}
	}
}
