package main

import (
	"github.com/eddielewis/pokedex/internal/pokeapi"
	"github.com/eddielewis/pokedex/internal/pokedex"
)

type config struct {
	pokeapiClient       pokeapi.Client
	nextLocationAreaURL *string
	prevLocationAreaURL *string
	pokedex             pokedex.Pokedex
}

func main() {
	cfg := config{
		pokeapiClient: pokeapi.NewClient(),
		pokedex:       pokedex.NewPokedex(),
	}

	startRepl(&cfg)
}
