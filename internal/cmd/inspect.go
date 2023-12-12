package cmd

import (
	"errors"
	"fmt"
)

func Inspect(cfg *Config) error {
	if len(cfg.Args) == 0 {
		return errors.New("No Pokemon Name Provided")
	}

	name := cfg.Args[0]
	pokemon, ok := cfg.Client.Pokedex[name]
	if !ok {
		return errors.New("Data Unavailable")
	}

	fmt.Fprintf(cfg.W, "Name: %s\n", pokemon.Name)
	fmt.Fprintf(cfg.W, "Height: %d\n", pokemon.Height)
	fmt.Fprintf(cfg.W, "Weight: %d\n", pokemon.Weight)

	fmt.Fprintf(cfg.W, "Stats:\n")
	for _, stat := range pokemon.Stats {
		fmt.Fprintf(cfg.W, "  - %s: %d\n", stat.Stat.Name, stat.BaseStat)
	}

	fmt.Fprintf(cfg.W, "Types:\n")
	for _, t := range pokemon.Types {
		fmt.Fprintf(cfg.W, "  - %s\n", t.Type.Name)
	}

	return nil
}
