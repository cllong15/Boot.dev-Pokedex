package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// GetPokemon -
func (c *Client) GetPokemon(pokemonName string) (Pokemon, error) {
	url := baseURL + "pokemon/" + pokemonName

	if val, ok := c.cache.Get(url); ok {
		pokemonResp := Pokemon{}
		err := json.Unmarshal(val, &pokemonResp)
		if err != nil {
			return Pokemon{}, fmt.Errorf("GetPokemon: cache Unmarshal: %v", err)
		}
		return pokemonResp, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return Pokemon{}, fmt.Errorf("GetPokemon: req: %v", err)
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return Pokemon{}, fmt.Errorf("GetPokemon: resp: %v", err)
	}
	defer resp.Body.Close()

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return Pokemon{}, fmt.Errorf("GetPokemon: dat: %v", err)
	}

	pokemonResp := Pokemon{}
	err = json.Unmarshal(dat, &pokemonResp)
	if err != nil {
		return Pokemon{}, fmt.Errorf("GetPokemon: pokeResp Unmarshal: %v", err)
	}

	c.cache.Add(url, dat)

	return pokemonResp, nil
}
