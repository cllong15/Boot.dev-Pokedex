package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// GetLocation -
func (c *Client) GetLocation(locationName string) (Location, error) {
	url := baseURL + "/location-area/" + locationName

	if val, ok := c.cache.Get(url); ok {
		locationResp := Location{}
		err := json.Unmarshal(val, &locationResp)
		if err != nil {
			return Location{}, fmt.Errorf("ListPokemon: cache Unmarshal: %v", err)
		}
		return locationResp, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return Location{}, fmt.Errorf("ListPokemon: req: %v", err)
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return Location{}, fmt.Errorf("ListPokemon: resp: %v", err)
	}
	defer resp.Body.Close()

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return Location{}, fmt.Errorf("ListPokemon: dat: %v", err)
	}

	locationResp := Location{}
	err = json.Unmarshal(dat, &locationResp)
	if err != nil {
		return Location{}, fmt.Errorf("ListPokemon: locationResp Unmarshal: %v", err)
	}

	c.cache.Add(url, dat)

	return locationResp, nil
}
