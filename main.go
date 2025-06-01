package main

import (
	"bufio"
	"fmt"
	"os"
)

type cliCommand struct {
	name        string
	description string
	callback    func()
}

func main() {
	var current Location
	err := current.location()
	if err != nil {
		fmt.Printf("Error from parsing current location: %v\n", err)
	}
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
		"map": {
			name:        "map",
			description: "Displays the current location",
			callback: func() {
				for _, value := range current.Results {
					fmt.Println(value.Name)
				}
			},
		},
		"nmap": {
			name:        "nmap",
			description: "Displays the next location",
			callback: func() {
				err := current.next()
				if err != nil {
					fmt.Printf("Error from parsing next location: %v\n", err)
				}
				for _, value := range current.Results {
					fmt.Println(value.Name)
				}
			},
		},

		"nmapb": {
			name:        "nmapb",
			description: "Displays the previous location",
			callback: func() {
				err := current.prev()
				if err != nil {
					fmt.Printf("Error from parsing previous location: %v\n", err)
				}
				for _, value := range current.Results {
					fmt.Println(value.Name)
				}
			},
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
			command.callback()
		}
	}
}
