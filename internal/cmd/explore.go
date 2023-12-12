package cmd

import (
	"errors"
	"fmt"
)

func Explore(cfg *Config) error {
	if len(cfg.Args) == 0 {
		return errors.New("No Location Area Provided")
	}

	area := cfg.Args[0]
	resp, err := cfg.Client.ExploreLocationArea(area)
	if err != nil {
		return err
	}

	fmt.Fprintf(cfg.W, "Exploring %s...\n\n", area)

	fmt.Fprintln(cfg.W, "Found Pokemon:")
	for _, encounter := range resp.PokemonEncounters {
		fmt.Fprintf(cfg.W, "- %s\n", encounter.Pokemon.Name)
	}

	return nil
}
