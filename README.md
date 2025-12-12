# Go Shell (Codecrafters)

A POSIX-inspired shell implemented in Go with command parsing, readline-powered REPL, builtin commands, history, redirection, and external process execution. 

[![demo](demo.gif)](demo.gif)


## Feature Highlights

### 🔹 Interactive REPL (Read-Eval-Print Loop)

The shell provides a modern, user-friendly interactive environment powered by the Readline library:

- **Dynamic Prompt Formatting**: Displays current working directory in the prompt for context
- **Autocomplete Support**: Builtin commands and shell functions are autocompleted as you type
- **Persistent Command History**: 
  - Commands are stored per-session in memory
  - History file location configurable via `HISTFILE` environment variable
  - Defaults to `/tmp/my_shell_history.tmp` if not specified
  - Automatic history flush on shell exit
- **Signal Handling**: Gracefully handles `Ctrl+C` (SIGINT) and `Ctrl+D` (EOF) without crashing
- **Line Editing**: Full support for arrow keys, backspace, delete, and other standard terminal editing features

### 🔹 Advanced Command Parsing

The lexer and parser components work together to intelligently interpret user input:

- **Tokenization**: Breaks input strings into logical tokens while respecting shell quoting rules
- **Abstract Syntax Tree (AST)**: Builds an AST representation of commands for validation and execution
- **Quote Handling**: Correctly processes single quotes (`'...'`), double quotes (`"..."`), and escape sequences
- **Error Detection**: Identifies syntax errors before attempting execution and provides clear error messages
- **Context Preservation**: Maintains command arguments, flags, and redirection targets throughout parsing

### 🔹 Built-in Command Suite

Nine essential builtin commands are implemented directly in Go for optimized performance:

- **File System Navigation**: `cd`, `pwd` for directory traversal
- **Text Output**: `echo` with proper argument handling
- **Shell Management**: `exit` with status codes, `history` for session review
- **Environment Control**: `set`, `unset`, `env` for variable management
- **Command Introspection**: `type` to determine if a command is builtin or external
- **In-Process Execution**: All builtins run directly in the shell process without spawning subprocesses

### 🔹 External Program Execution

The executor seamlessly launches external programs while maintaining shell environment consistency:

- **Process Spawning**: Creates child processes for external commands (e.g., `ls`, `grep`, `cat`)
- **Environment Inheritance**: External programs inherit all current environment variables
- **Argument Passing**: Properly forwards all command arguments and flags to external executables
- **Standard Stream Integration**: Handles stdin, stdout, and stderr from spawned processes
- **Exit Code Propagation**: Captures and reports exit codes from external programs

### 🔹 Input/Output Redirection

Full support for redirecting output streams to files, with intelligent directory creation:

- **Stdout Redirection**:
  - `> file` - Write stdout to file (overwrite)
  - `>> file` - Append stdout to file
  - `1> file` / `1>> file` - Explicit stdout redirection
- **Stderr Redirection**:
  - `2> file` - Write stderr to file (overwrite)
  - `2>> file` - Append stderr to file
- **Descriptor-Prefixed Variants**: Supports formats like `2>>logfile` (no spaces)
- **Automatic Directory Creation**: Creates parent directories as needed (equivalent to `mkdir -p`)
- **File Mode Handling**: Creates files with appropriate permissions (0644 for regular files)
- **Error Handling**: Gracefully handles file permission errors and reports them to the user

### 🔹 Command History Management

Tracks all executed commands throughout the session with persistent storage:

- **Session History**: In-memory list of all commands executed in the current session
- **History Command**: `history` builtin displays all past commands with line numbers
- **Persistent Storage**: History is saved to the file specified by `HISTFILE` when shell exits
- **Automatic Flushing**: Ensures no command history is lost on unexpected termination
- **History Recall**: Previous history entries can be navigated using readline's history features

### 🔹 Extensible Architecture

The codebase is thoughtfully organized to make adding new features straightforward:

- **Modular Design**: Separate packages for parsing, execution, and REPL interaction
- **Clear Interfaces**: Each component (lexer, parser, executor) has well-defined responsibilities
- **Easy Feature Addition**: Adding pipes (`|`), globbing (`*`, `?`), or new builtins requires minimal changes
- **Decoupled Layers**: Parser doesn't depend on executor; both are independent of REPL
- **Testability**: Modular structure makes unit testing individual components straightforward

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
cd /tmp
echo "hello" > out.txt
echo append >> out.txt
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
