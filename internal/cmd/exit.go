package cmd

import "os"

func Exit(_ Config) error {
	os.Exit(0)
	return nil
}
