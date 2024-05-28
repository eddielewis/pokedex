package main

import (
	"fmt"
)

func callbackInspect(cfg *config, args ...string) error {
	name := args[0]

	pokemon, ok := cfg.pokedex.Get(name)
	if !ok {
		return fmt.Errorf("you have not caught that pokemon yet")
	}

	fmt.Printf("Name: %s\n", pokemon.Name)
	fmt.Printf("Height: %d\n", pokemon.Height)
	fmt.Printf("Weight: %d\n", pokemon.Weight)
	fmt.Printf("Stats:\n")
	for _, stat := range pokemon.Stats {
		fmt.Printf(" -%s: %d\n", stat.Stat.Name, stat.BaseStat)
	}
	fmt.Printf("Types:\n")
	for _, pType := range pokemon.Types {
		fmt.Printf(" -%s\n", pType.Type.Name)
	}

	fmt.Println("")
	return nil
}
