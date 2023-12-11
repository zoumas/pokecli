package main

import (
	"net/http"
	"os"
	"time"

	"github.com/zoumas/pokecli/internal/cmd"
	"github.com/zoumas/pokecli/internal/pokeapi"
	"github.com/zoumas/pokecli/internal/repl"
)

func main() {
	// TODO: Look into the functional options struct configuration pattern
	repl.Start(
		repl.NewConfig(
			"pokecli > ",
			os.Stdin,
			os.Stdout,
			pokeapi.NewClient(&http.Client{Timeout: 30 * time.Second}),
			cmd.Cmds(),
		),
	)
}
