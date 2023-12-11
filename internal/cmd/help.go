package cmd

import (
	"fmt"
	"io"
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
	fmt.Fprint(w, "\nCommands:\n\n")

	for _, c := range cmds {
		fmt.Fprintf(w, "%s - %-4s\n", c.Name, c.Description)
	}
	fmt.Fprintln(w)
}
