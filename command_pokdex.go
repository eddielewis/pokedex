package main

import (
	"fmt"
)

func callbackPokedex(cfg *config, args ...string) error {
	if len(cfg.pokedex.Pokedex) < 1 {
		return fmt.Errorf("you haven't caught any Pokemon yet")
	}

	fmt.Printf("Your Pokedex:\n")
	for key, _ := range cfg.pokedex.Pokedex {
		fmt.Printf(" -%s\n", key)
	}

	fmt.Println("")
	return nil
}
