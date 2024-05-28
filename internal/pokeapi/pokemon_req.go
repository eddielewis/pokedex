package pokeapi

import "encoding/json"

func (c *Client) CatchPokemon(pokemonName string) (Pokemon, error) {
	url := baseURL + "/pokemon/" + pokemonName

	data, ok := c.cache.Get(url)
	if !ok {
		var err error
		data, err = c.MakeRequest(url)
		if err != nil {
			return Pokemon{}, err
		}
	}

	pokemonResp := Pokemon{}
	err := json.Unmarshal(data, &pokemonResp)
	if err != nil {
		return Pokemon{}, err
	}

	c.cache.Add(url, data)

	return pokemonResp, nil
}
