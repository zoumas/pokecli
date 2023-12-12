package cmd

import (
	"errors"
	"fmt"
	"math/rand"
)

func Catch(cfg *Config) error {
	if len(cfg.Args) == 0 {
		return errors.New("No Pokemon Name Provided")
	}

	name := cfg.Args[0]
	pokemon, err := cfg.Client.GetPokemon(name)
	if err != nil {
		return err
	}

	fmt.Fprintf(cfg.W, "Throwing a Pokeball at %s\n", name)

	const threshold = 50
	catchRate := rand.Intn(pokemon.BaseExperience)
	if catchRate > threshold {
		fmt.Fprintf(cfg.W, "%s escaped!\n", name)
		return nil
	}

	fmt.Fprintf(cfg.W, "%s was caught!\n", name)
	cfg.Client.AddToPokedex(name, pokemon)

	return nil
}
