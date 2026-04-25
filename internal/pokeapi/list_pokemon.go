package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// ListLocations -
func (c *Client) ListPokemon(url *string) ([]PokemonEncounters, error) {
	// fmt.Println(*url)
	// list pokemon from area
	if val, ok := c.cache.Get(*url); ok {
		locationsResp := RespLocationAreas{}
		err := json.Unmarshal(val, &locationsResp)
		if err != nil {
			return []PokemonEncounters{}, fmt.Errorf("ListPokemon: cache Unmarshal: %v", err)
		}

		return locationsResp.PokemonEncounters, nil
	}

	req, err := http.NewRequest("GET", *url, nil)
	if err != nil {
		return []PokemonEncounters{}, fmt.Errorf("ListPokemon: req: %v", err)
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return []PokemonEncounters{}, fmt.Errorf("ListPokemon: resp: %v", err)
	}
	defer resp.Body.Close()

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return []PokemonEncounters{}, fmt.Errorf("ListPokemon: dat: %v", err)
	}

	locationsResp := RespLocationAreas{}
	err = json.Unmarshal(dat, &locationsResp)
	if err != nil {
		return []PokemonEncounters{}, fmt.Errorf("ListPokemon: locationResp Unmarshal: %v", err)
	}

	c.cache.Add(*url, dat)
	return locationsResp.PokemonEncounters, nil
}
