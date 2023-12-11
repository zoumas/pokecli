package main

import (
	"os"

	"github.com/zoumas/pokecli/internal/cmd"
	"github.com/zoumas/pokecli/internal/repl"
)

const prompt = "pokecli > "

func main() {
	repl.Start(repl.NewConfig(prompt, os.Stdin, os.Stdout, cmd.Cmds()))
}
