package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func cleanInput(text string) []string {
	ouput := strings.ToLower(text)
	words := strings.Fields(ouput) 
	return words
}

func startRepl() {
	reader := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")

		if err := reader.Err(); err != nil {
			log.Fatal("Scan error: ", err)
		}
		reader.Scan()

		words := cleanInput(reader.Text())
		if len(words) == 0 {
			continue
		} 

		commandName := words[0]
		fmt.Printf("Your command was: %v\n", commandName)
	}
}
