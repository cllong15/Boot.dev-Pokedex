package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func commandMap(config *config) error {
	// get request for next page
	if config.Next == "" {
		fmt.Print("No next page\n")
		return nil
	}
	res, ok := http.Get(config.Next)
	if ok != nil {
		return ok
	}
	defer res.Body.Close()

	var areas AreaStruct
	decoder := json.NewDecoder(res.Body)
	ok = decoder.Decode(&areas)
	if ok != nil {
		return ok
	}
	if areas.Next != nil {
		config.Next = *areas.Next
	}
	if areas.Previous != nil {
		config.Previous = *areas.Previous
	}
	for _, area := range areas.Results {
		fmt.Println(area.Name)
	}
	return nil
}
