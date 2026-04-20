package main

import (
	"fmt"
	"strings"
	"bufio"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for ;; {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		input := scanner.Text()
		output := CleanInput(input)[0]
		fmt.Printf("Your command was: %v\n", output)
	}
}

func CleanInput(text string) []string {
	text = strings.ToLower(text)
	text = strings.TrimSpace(text)
	return strings.SplitN(text, " ", 2)
}