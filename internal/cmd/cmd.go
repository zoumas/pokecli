package cmd

import "io"

// Config is a command's configuration.
// It contains its arguments and any additional data a command may need.
type Config struct {
	W    io.Writer
	Args []string
}

// Cmd represents a REPL command.
type Cmd struct {
	Name        string
	Description string
	Callback    func(Config) error
}

// Cmds returns a map of all the available commands.
func Cmds() map[string]Cmd {
	return map[string]Cmd{
		"help": {
			Name:        "help",
			Description: "Display a usage menu or get information about a certain command",
			Callback:    Help,
		},
		"exit": {
			Name:        "exit",
			Description: "Exit pokecli",
			Callback:    Exit,
		},
	}
}
