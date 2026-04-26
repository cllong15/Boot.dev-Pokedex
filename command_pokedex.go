package main

import "fmt"

func commandPokedex(cfg *config, args ...string) error {
	fmt.Println("Your Pokedex:")
	for _, poke := range cfg.caughtPokemon {
		fmt.Printf(" - %s\n", poke.Name)
	}
	return nil
}
