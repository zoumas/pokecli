package cmd

import (
	"io"

	"github.com/zoumas/pokecli/internal/pokeapi"
)

// Config is a command's configuration.
// It contains its arguments and any additional data a command may need.
type Config struct {
	Args   []string
	W      io.Writer
	Client *pokeapi.Client
}
