package cmd

import (
	"fmt"

	"github.com/zoumas/pokecli/internal/pokeapi"
)

func Map(cfg *Config) error {
	url := pokeapi.LocationAreaEndpoint
	if cfg.Client.Next != nil {
		url = *cfg.Client.Next
	}

	resp, err := cfg.Client.GetLocationAreas(url)
	if err != nil {
		return err
	}

	fmt.Fprintln(cfg.W)
	for _, r := range resp.Results {
		fmt.Fprintln(cfg.W, r.Name)
	}
	fmt.Fprintln(cfg.W)

	cfg.Client.Next = resp.Next
	cfg.Client.Previous = resp.Previous

	return nil
}
