package cmd

// Cmd represents a REPL command.
type Cmd struct {
	Name        string
	Description string
	Callback    func(*Config) error
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
		"map": {
			Name:        "map",
			Description: "Explore the next 20 areas in the Pokemon world",
			Callback:    Map,
		},
		"mapb": {
			Name:        "mapb",
			Description: "Explore the previous 20 areas in the Pokemon world",
			Callback:    Mapb,
		},
		"explore": {
			Name:        "explore {area}",
			Description: "Get a list of possible pokemon encounters for the provided area",
			Callback:    Explore,
		},
	}
}
