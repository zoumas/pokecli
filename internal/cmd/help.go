package cmd

import (
	"cmp"
	"fmt"
	"io"
	"slices"
)

type CmdErr string

func (e CmdErr) Error() string {
	return string(e)
}

func ErrUnknown(name string) CmdErr {
	return CmdErr("unknown help topic: " + name)
}

func Help(cfg *Config) error {
	cmds := Cmds()

	if len(cfg.Args) == 0 {
		helpMenu(cfg.W, cmds)
		return nil
	}

	name := cfg.Args[0]
	c, ok := cmds[name]
	if !ok {
		return ErrUnknown(name)
	}

	fmt.Fprintf(cfg.W, "%s: %s\n", name, c.Description)

	return nil
}

func helpMenu(w io.Writer, cmds map[string]Cmd) {
	fmt.Fprintln(w, "Commands:")

	sortedCommands := []Cmd{}
	for _, c := range cmds {
		sortedCommands = append(sortedCommands, c)
	}
	slices.SortStableFunc(sortedCommands, func(a, b Cmd) int {
		return cmp.Compare(a.Name, b.Name)
	})

	for _, c := range sortedCommands {
		fmt.Fprintf(w, "\n%s\n %s\n", c.Name, c.Description)
	}
}
