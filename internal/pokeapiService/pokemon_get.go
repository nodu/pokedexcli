package pokeapiService

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) GetPokemon(name string) (PokemonResonse, error) {
	fullURL := baseURL + "/pokemon/" + name

	cachedData, found := c.cache.Get(fullURL)
	if found {
		pokemon := PokemonResonse{}
		err := json.Unmarshal(cachedData, &pokemon)
		if err != nil {
			return PokemonResonse{}, err
		}

		return pokemon, nil
	}

	req, err := http.NewRequest("GET", fullURL, nil)

	if err != nil {
		return PokemonResonse{}, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return PokemonResonse{}, err
	}
	defer res.Body.Close()

	if res.StatusCode > 399 {
		return PokemonResonse{}, fmt.Errorf("Bad status code: %v", res.StatusCode)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return PokemonResonse{}, err
	}

	pokemon := PokemonResonse{}
	err = json.Unmarshal(body, &pokemon)

	if err != nil {
		return PokemonResonse{}, err
	}
	c.cache.Add(fullURL, body)
	return pokemon, nil
}
