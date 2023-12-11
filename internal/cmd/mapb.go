package cmd

import (
	"fmt"
)

func Mapb(cfg *Config) error {
	if cfg.Client.Previous == nil {
		fmt.Fprintln(cfg.W, "There is nothing back there")
		return nil
	}

	resp, err := cfg.Client.GetLocationAreas(*cfg.Client.Previous)
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
