package repl

import (
	"bufio"
	"fmt"
	"io"

	"github.com/zoumas/pokecli/internal/cmd"
)

// Start starts a Read-Eval-Print-Loop.
// Getting input from the given 'in' io.Reader and executing the corresponding commands based on a 'cmds' map of available commands.
func Start(cfg *Config) {
	scanner := bufio.NewScanner(cfg.r)
	for fmt.Print(cfg.prompt); scanner.Scan(); fmt.Print(cfg.prompt) {
		text, err := scanner.Text(), scanner.Err()
		if err != nil {
			fmt.Fprintln(cfg.w, "Error while reading input:", err)
			continue
		}

		input := CleanInput(text)
		if len(input) == 0 {
			continue
		}

		name := input[0]
		args := input[1:]

		c, ok := cfg.cmds[name]
		if !ok {
			handleUnknownCommand(cfg.w, name)
			continue
		}

		if err := c.Callback(&cmd.Config{
			Args:   args,
			W:      cfg.w,
			Client: cfg.client,
		}); err != nil {
			fmt.Fprintln(cfg.w, err)
		}
	}
}

func handleUnknownCommand(w io.Writer, name string) {
	fmt.Fprintf(
		w,
		"\nUnknown command '%s'\nTry 'help' to get a list of all the available commands\n\n",
		name,
	)
}
