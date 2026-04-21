package main

import (
	"fmt"
	"maps"
	"os"
	"slices"
)

type CLICommand struct {
	name        string
	description string
	callback    func() error
}

func getCommands() map[string]CLICommand {
	return map[string]CLICommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
	}
}

func commandExit() error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp() error {
	fmt.Print("Welcome to the Pokedex!\nUsage:\n\n")
	for i, command := range slices.Sorted(maps.Keys(getCommands())) {
		fmt.Printf("%d. %s: %s\n", i, getCommands()[command].name, getCommands()[command].description)
	}
	return nil
}
