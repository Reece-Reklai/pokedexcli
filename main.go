package main

import (
	"bufio"
	"fmt"
	"github.com/Reece-Reklai/pokedexcli/internal/location"
	"github.com/Reece-Reklai/pokedexcli/test"
	"os"
)

type cliCommand struct {
	name        string
	description string
	callback    func()
}

func main() {
	var current location.Location
	err := current.Location()
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
		"nmapb": {
			name:        "nmapb",
			description: "Displays the previous location",
			callback: func() {
				err := current.Prev_map()
				if err != nil {
					fmt.Printf("Error from parsing previous location: %v\n", err)
				}
				for _, value := range current.Results {
					fmt.Println(value.Name)
				}
			},
		},
		"nmap": {
			name:        "nmap",
			description: "Displays the next location",
			callback: func() {
				err := current.Next_map()
				if err != nil {
					fmt.Printf("Error from parsing next location: %v\n", err)
				}
				for _, value := range current.Results {
					fmt.Println(value.Name)
				}
			},
		},
	}
	scanned_input := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex >>>>>>>>>>>>>>>>> ")
		scanned_input.Scan()
		text := scanned_input.Text()
		clean := test.CleanInput(text)
		if clean == nil {
			continue
		}
		command, ok := cli[clean[0]]
		if ok == false {
			fmt.Println("Unknown Command")
		} else {
			switch command.name {
			case "exit":
				command.callback()
			case "help":
				command.callback()
			case "map":
				command.callback()
			case "nmapb":
				if current.Previous == nil {
					fmt.Println("There are no known locations in the previous map")
					continue
				}
				command.callback()
			case "nmap":
				if current.Next == nil {
					fmt.Println("There are no known locations in the next map")
					continue
				}
				command.callback()
			}
		}
	}
}
