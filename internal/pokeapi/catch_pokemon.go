package pokeapi

import (
	"io"
	"net/http"
	"encoding/json"
)

func (c *Client) CatchPokemon(pokemonName string) (Pokemon, error) {
	url := baseURL + "/pokemon/" + pokemonName

	var responseData []byte
	responseData, ok := c.cache.Get(url)
	if !ok {
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return Pokemon{}, err
		}

		resp, err := c.httpClient.Do(req)
		if err != nil {
			return Pokemon{}, err
		}
		defer resp.Body.Close()

		data, err := io.ReadAll(resp.Body)
		if err != nil {
			return Pokemon{}, err
		}

		c.cache.Add(url, data)

		responseData = data
	}

	var pokemon Pokemon
	err := json.Unmarshal(responseData, &pokemon)
	if err != nil {
		return Pokemon{}, err
	}

	return pokemon, nil
}
