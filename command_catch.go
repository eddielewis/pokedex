package main

import (
	"fmt"
	"math/rand"
)

func callbackCatch(cfg *config, args ...string) error {
	pokemonName := args[0]
	fmt.Printf("Throwing a Pokeball at %s...\n", pokemonName)

	pokemon, err := cfg.pokeapiClient.CatchPokemon(pokemonName)
	if err != nil {
		return err
	}

	const threshold = 50
	randNum := rand.Intn(pokemon.BaseExperience)
	if randNum > threshold {
		return fmt.Errorf("%s escaped", pokemonName)
	}

	cfg.pokedex.Add(pokemon)
	fmt.Printf("%s was caught!\n", pokemonName)

	return nil
}
