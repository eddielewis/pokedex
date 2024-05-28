package pokedex

import (
	"sync"

	"github.com/eddielewis/pokedex/internal/pokeapi"
)

type Pokedex struct {
	Pokedex map[string]pokeapi.Pokemon
	mux     *sync.Mutex
}

func NewPokedex() Pokedex {
	return Pokedex{
		Pokedex: make(map[string]pokeapi.Pokemon),
		mux:     &sync.Mutex{},
	}
}

func (p *Pokedex) Add(pokemon pokeapi.Pokemon) {
	p.mux.Lock()
	defer p.mux.Unlock()
	p.Pokedex[pokemon.Name] = pokemon
}

func (p *Pokedex) Get(name string) (pokeapi.Pokemon, bool) {
	p.mux.Lock()
	defer p.mux.Unlock()
	pokemon, ok := p.Pokedex[name]
	return pokemon, ok
}
