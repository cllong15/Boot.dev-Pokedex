package main

import "fmt"

func commandInspect(cfg *config, args ...string) error {
	pokemon, ok := cfg.caughtPokemon[args[0]]
	if ok == false {
		return fmt.Errorf("you have not caught that pokemon")
	}

	fmt.Printf("Name: %s\n", pokemon.Name)
	fmt.Printf("Height: %d\n", pokemon.Height)
	fmt.Printf("Weight: %d\n", pokemon.Weight)
	fmt.Printf("Stats:\n")
	fmt.Printf(" -hp: %d\n", pokemon.Stats[0].BaseStat)
	fmt.Printf(" -attack: %d\n", pokemon.Stats[1].BaseStat)
	fmt.Printf(" -defense: %d\n", pokemon.Stats[2].BaseStat)
	fmt.Printf(" -special-attack: %d\n", pokemon.Stats[3].BaseStat)
	fmt.Printf(" -special-defense: %d\n", pokemon.Stats[4].BaseStat)
	fmt.Printf(" -speed: %d\n", pokemon.Stats[5].BaseStat)
	fmt.Printf("Types:\n")
	fmt.Printf(" - %s\n", pokemon.Types[0].Type.Name)
	if len(pokemon.Types) > 1 {
		fmt.Printf(" - %s\n", pokemon.Types[1].Type.Name)
	}
	return nil
}
