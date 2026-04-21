package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startRepl() {
	reader := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		reader.Scan()

		words := cleanInput(reader.Text())
		if len(words) == 0 {
			continue
		}

		var cliCommands = getCommands()
		var command CLICommand
		var ok bool
		command, ok = cliCommands[words[0]]

		if ok == false {
			fmt.Print("Error: Command not found\n")
		} else {
			command.callback()
		}
	}
}

func cleanInput(text string) []string {
	output := strings.ToLower(text)
	words := strings.Fields(output)
	return words
}
