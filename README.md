# Go Shell (Codecrafters)

A POSIX-inspired shell implemented in Go with command parsing, readline-powered REPL, builtin commands, history, redirection, and external process execution. 

## Feature Highlights

| Area | Details |
| --- | --- |
| REPL | Readline prompt with autocomplete for builtins, persistent history via `HISTFILE` (defaults to `/tmp/my_shell_history.tmp`), and graceful EOF/interrupt handling. |
| Parsing | Tokenizes input into an AST, supports quoting, and detects syntax errors before execution. |
| Builtins | `cd`, `pwd`, `echo`, `history`, `set`, `unset`, `env`, `type`, `exit` implemented in Go for fast, in-process execution. |
| External execution | Spawns external programs with inherited env and proper arg handling. |
| Redirection | Supports `>`, `>>`, `1>`, `2>`, and descriptor-prefixed variants (e.g. `2>>log`). Creates parent directories when needed. |
| History | Tracks commands per-session and flushes to history file on exit. |
| Extensibility | Clear separation between parsing, execution, and REPL layers makes it easy to add pipes, globbing, or more builtins. |

### Builtin Commands

| Command | Behavior |
| --- | --- |
| `cd [path]` | Change directory; defaults to `$HOME` when no path provided. |
| `pwd` | Print current working directory. |
| `echo [args...]` | Print arguments with newline handling. |
| `history` | Show in-memory history for the current session. |
| `set` / `unset` | Manage environment variables in the running shell. |
| `env` | List current environment variables. |
| `type <name>` | Show whether a token is a builtin or an external command. |
| `exit [status]` | Terminate the shell with an optional status code. |

## Project Layout

```
.
├─ app/
│  └─ main.go            # Entry point wiring the REPL and executor
├─ internal/
│  ├─ exec/              # Builtins, external execution, redirection, history
│  ├─ parser/            # Lexer, AST, and parser for shell input
│  └─ repl/              # Prompt, autocomplete, history integration
├─ codecrafters.yml      # Challenge config
├─ go.mod                # Go module definition
├─ README.md             # You are here
└─ your_program.sh       # Helper runner for the challenge harness
```

## Getting Started

1. Prerequisites: Go `1.25` or newer.
2. Install dependencies (none beyond the standard library and bundled modules):
   ```sh
   go mod tidy
   ```
3. Run the shell locally:
   ```sh
   go run ./app
   ```

## Usage Examples

```sh
pwd
cd /tmp && echo "hello" > out.txt
echo append >> out.txt
env | grep HOME
type cd
history
exit 0
```

## Development Notes

- Prompt styling and history location can be customized via environment variables (`HISTFILE` for persistence).
- Redirection handles stdout/stderr independently and will create missing parent directories.
- Parsing, execution, and REPL layers are intentionally decoupled to make adding pipes or glob expansion straightforward.

## Credits

Built as part of the [Codecrafters](https://codecrafters.io) platform. Huge thanks to their course for the framework and staged challenges that guide the implementation. This repository is a solution for the [Codecrafters "Build Your Own Shell" course](https://app.codecrafters.io/courses/shell/overview).
