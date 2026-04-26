package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// GetLocation -
func (c *Client) GetPokemon(PokeName string) (Pokemon, error) {
	url := baseURL + "pokemon/" + PokeName
	// fmt.Println(url)

	if val, ok := c.cache.Get(url); ok {
		pokeResp := Pokemon{}
		err := json.Unmarshal(val, &pokeResp)
		if err != nil {
			return Pokemon{}, fmt.Errorf("GetPokemon: cache Unmarshal: %v", err)
		}
		return pokeResp, nil
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

	pokeResp := Pokemon{}
	err = json.Unmarshal(dat, &pokeResp)
	if err != nil {
		return Pokemon{}, fmt.Errorf("GetPokemon: pokeResp Unmarshal: %v", err)
	}

	c.cache.Add(url, dat)

	return pokeResp, nil
}
