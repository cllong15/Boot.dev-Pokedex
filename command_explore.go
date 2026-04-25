package main

import "fmt"

func commandExplore(cfg *config, loc_name *string) error {
	if *loc_name == "" {
		return fmt.Errorf("No location name")
	}
	url := "https://pokeapi.co/api/v2/location-area/" + *loc_name
	encounters, err := cfg.pokeapiClient.ListPokemon(&url)
	if err != nil {
		return fmt.Errorf("commandExplore encounters %v", err)
	}
	for _, enc := range encounters {
		fmt.Println(enc.Pokemon.Name)
	}
	return nil

}
