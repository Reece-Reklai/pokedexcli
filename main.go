package main

import (
	"bufio"
	"fmt"
	"os"
)

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

func main() {
	cli := map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Displays a help messsage",
			callback:    commandHelp,
		},
	}
	scanned_input := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanned_input.Scan()
		text := scanned_input.Text()
		clean := CleanInput(text)
		if clean == nil {
			continue
		}
		command, ok := cli[clean[0]]
		if ok == false {
			fmt.Println("Unknown Command")
		} else {
			err := command.callback()
			if err != nil {
				fmt.Printf("Command: %s failed", command.name)
			}
		}
	}
}
