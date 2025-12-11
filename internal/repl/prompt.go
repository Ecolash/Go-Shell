package repl

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

// Basic Colors
const (
	RESET   = "\033[0m"
	CYAN    = "\033[36m"
	GREEN   = "\033[32m"
	YELLOW  = "\033[33m"
	RED     = "\033[31m"
	MAGENTA = "\033[35m"
)

// Foreground colors
const (
	FG_DEFAULT = "\033[39m"
	FG_BLACK   = "\033[30m"
	FG_RED     = "\033[31m"
	FG_GREEN   = "\033[32m"
	FG_YELLOW  = "\033[33m"
	FG_BLUE    = "\033[34m"
	FG_MAGENTA = "\033[35m"
	FG_CYAN    = "\033[36m"
	FG_WHITE   = "\033[37m"

	// Bright (bold) versions
	FG_BRIGHT_BLACK   = "\033[90m"
	FG_BRIGHT_RED     = "\033[91m"
	FG_BRIGHT_GREEN   = "\033[92m"
	FG_BRIGHT_YELLOW  = "\033[93m"
	FG_BRIGHT_BLUE    = "\033[94m"
	FG_BRIGHT_MAGENTA = "\033[95m"
	FG_BRIGHT_CYAN    = "\033[96m"
	FG_BRIGHT_WHITE   = "\033[97m"
)

// Background colors
const (
	BG_DEFAULT = "\033[49m"
	BG_BLACK   = "\033[40m"
	BG_RED     = "\033[41m"
	BG_GREEN   = "\033[42m"
	BG_YELLOW  = "\033[43m"
	BG_BLUE    = "\033[44m"
	BG_MAGENTA = "\033[45m"
	BG_CYAN    = "\033[46m"
	BG_WHITE   = "\033[47m"

	// Bright (bold) backgrounds
	BG_BRIGHT_BLACK   = "\033[100m"
	BG_BRIGHT_RED     = "\033[101m"
	BG_BRIGHT_GREEN   = "\033[102m"
	BG_BRIGHT_YELLOW  = "\033[103m"
	BG_BRIGHT_BLUE    = "\033[104m"
	BG_BRIGHT_MAGENTA = "\033[105m"
	BG_BRIGHT_CYAN    = "\033[106m"
	BG_BRIGHT_WHITE   = "\033[107m"
)

func getGitInfo(dir string) (branch string, dirty bool) {
	if _, err := os.Stat(filepath.Join(dir, ".git")); err != nil {
		return "", false
	}

	headBytes, err := os.ReadFile(filepath.Join(dir, ".git/HEAD"))
	if err != nil {
		return "", false
	}

	head := strings.TrimSpace(string(headBytes))
	if strings.HasPrefix(head, "ref: ") {
		parts := strings.Split(head[5:], "/")
		branch = parts[len(parts)-1]
	} else {
		branch = head[:7]
	}
	cmd := exec.Command("git", "-C", dir, "status", "--porcelain")
	out, err := cmd.Output()
	if err != nil {
		return branch, false
	}
	if len(out) > 0 {
		dirty = true
	}
	return branch, dirty
}

func formatPrompt() string {
	cwd, err := os.Getwd()
	if err != nil {
		cwd = "?"
	}

	home := os.Getenv("HOME")
	if home != "" && strings.HasPrefix(cwd, home) {
		cwd = "~" + strings.TrimPrefix(cwd, home)
	}

	// Icons (require Nerd Font; harmless if unsupported)
	folderIcon := "" // nf-fa-folder
	gitIcon := " "   // nf-dev-git_branch
	arrowIcon := ">"

	branch, dirty := getGitInfo(".")
	gitPart := ""
	if branch != "" {
		if dirty {
			gitPart = fmt.Sprintf(" %s%s*%s", gitIcon, YELLOW+branch, RESET)
		} else {
			gitPart = fmt.Sprintf(" %s%s%s", gitIcon, GREEN+branch, RESET)
		}
	}

	// Final colored prompt
	return fmt.Sprintf(
		"%s%s %s%s%s %s ",
		CYAN, folderIcon, YELLOW, cwd, RESET,
		gitPart+MAGENTA+" "+arrowIcon+RESET,
	)
}
