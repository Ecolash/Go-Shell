package parser

type Command struct {
	Name   string
	Args   []string
	PipeTo *Command
}

type Pipeline struct {
	Commands []*Command
}
