package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// ListLocations -
func (c *Client) ListLocations(pageURL *string) (RespShallowLocations, error) {
	url := baseURL + "/location-area"
	if pageURL != nil {
		url = *pageURL
	}

	if val, ok := c.cache.Get(url); ok {
		locationsResp := RespShallowLocations{}
		err := json.Unmarshal(val, &locationsResp)
		if err != nil {
			return RespShallowLocations{}, fmt.Errorf("ListLocations: cache Unmarshal: %v", err)
		}

		return locationsResp, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return RespShallowLocations{}, fmt.Errorf("ListLocations: req: %v", err)
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return RespShallowLocations{}, fmt.Errorf("ListLocations: resp: %v", err)
	}
	defer resp.Body.Close()

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return RespShallowLocations{}, fmt.Errorf("ListLocations: dat: %v", err)
	}

	locationsResp := RespShallowLocations{}
	err = json.Unmarshal(dat, &locationsResp)
	if err != nil {
		return RespShallowLocations{}, fmt.Errorf("ListLocations: locationResp Unmarshal: %v", err)
	}

	c.cache.Add(url, dat)
	return locationsResp, nil
}
