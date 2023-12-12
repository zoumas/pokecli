package cmd

import (
	"cmp"
	"fmt"
	"slices"
)

type PokedexEntry struct {
	ID   int
	Name string
}

func Pokedex(cfg *Config) error {
	if len(cfg.Client.Pokedex) == 0 {
		return nil
	}

	pokedex := []PokedexEntry{}
	for k, v := range cfg.Client.Pokedex {
		pokedex = append(pokedex, PokedexEntry{
			ID:   v.ID,
			Name: k,
		})
	}

	slices.SortStableFunc(pokedex, func(a, b PokedexEntry) int {
		return cmp.Compare(a.ID, b.ID)
	})

	for _, p := range pokedex {
		fmt.Fprintf(cfg.W, "%3d - %s\n", p.ID, p.Name)
	}

	return nil
}
