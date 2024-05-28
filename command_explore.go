package main

import (
	"fmt"
)

func callbackExplore(cfg *config, args ...string) error {
	areaName := args[0]
	resp, err := cfg.pokeapiClient.ExploreLocationArea(areaName)
	if err != nil {
		return err
	}
	fmt.Println("Possible pokemon:")
	for _, area := range resp.PokemonEncounters {
		fmt.Printf("- %s\n", area.Pokemon.Name)
	}

	fmt.Println("")
	return nil
}
