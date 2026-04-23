package main

import "fmt"

func commandHelp(config *config) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	for _, cmd := range getCommands() {
		fmt.Printf("\n%s: %s", cmd.name, cmd.description)
	}
	fmt.Println()
	return nil
}
