package repl

import (
	"io"

	"github.com/zoumas/pokecli/internal/cmd"
	"github.com/zoumas/pokecli/internal/pokeapi"
)

// Config stores any data the REPL may need in order to operate.
type Config struct {
	prompt string
	r      io.Reader
	w      io.Writer
	client *pokeapi.Client
	cmds   map[string]cmd.Cmd
}

// NewConfig constructs a configuration for the REPL
func NewConfig(
	prompt string,
	r io.Reader,
	w io.Writer,
	client *pokeapi.Client,
	cmds map[string]cmd.Cmd,
) *Config {
	return &Config{
		prompt: prompt,
		r:      r,
		w:      w,
		client: client,
		cmds:   cmds,
	}
}
