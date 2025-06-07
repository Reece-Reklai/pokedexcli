package main

import (
	"bufio"
	"fmt"
	"os"
	"time"

	"github.com/Reece-Reklai/pokedexcli/internal/catch"
	"github.com/Reece-Reklai/pokedexcli/internal/explore"
	"github.com/Reece-Reklai/pokedexcli/internal/location"
	"github.com/Reece-Reklai/pokedexcli/internal/pokecache"
	"github.com/Reece-Reklai/pokedexcli/test"
)

type cliCommand struct {
	name        string
	description string
	callback    func()
}

func main() {
	var current location.Location
	var encounter explore.Explore
	var capture catch.Pokemon
	var locationArea string
	var pokemon string
	err := current.Location()
	if err != nil {
		fmt.Printf("Error from parsing current location: %v\n", err)
	}
	locationCache := pokecache.NewLocation(time.Second)
	encounterCache := pokecache.NewEncounter(time.Second)
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
		"catch": {
			name:        "catch",
			description: "Attempt to catch pokemon",
			callback: func() {
				err := capture.Catch(pokemon)
				if err != nil {
					fmt.Println(err)
				}
				fmt.Println(capture.Name)
			},
		},
		"explore": {
			name:        "explore",
			description: "Pokemon encountered in location-area",
			callback: func() {
				encounters, ok := encounterCache.EncounterGet(locationArea)
				if ok == true {
					fmt.Printf("Exploring %v...\n", locationArea)
					for _, value := range encounter.PokemonEncounters {
						fmt.Println(value.Pokemon.Name)
					}
				} else {
					found := false
					for _, value := range current.Results {
						if locationArea == value.Name {
							err := encounter.Explore(value.URL)
							if err != nil {
								fmt.Println(err)
							} else {
								encounterCache.EncounterAdd(locationArea, encounters)
							}
							found = true
							break
						}
					}
					if found == true {
						fmt.Printf("Exploring %v...\n", locationArea)
						for _, value := range encounter.PokemonEncounters {
							fmt.Println(value.Pokemon.Name)
						}
					} else {
						fmt.Println("Location-area Invalid")
					}
				}
			},
		},
		"nmapb": {
			name:        "nmapb",
			description: "Displays the previous location",
			callback: func() {
				if current.Previous == nil {
					fmt.Println("There are no known locations in the previous map")
				} else {
					previous, bool := locationCache.LocationGet(*current.Previous)
					for _, val := range previous {
						fmt.Println(val.Name)
					}
					if bool == true {
						url, err := current.PrevMap()
						resultsCopy := make([]struct {
							Name string `json:"name"`
							URL  string `json:"url"`
						}, len(current.Results))
						copy(resultsCopy, current.Results)
						if err != nil {
							fmt.Println(err)
						}
						if url != nil {
							locationCache.LocationAdd(*url, resultsCopy)
						} else {
							fmt.Println("Could not catch the previous previous map locations")
						}
						for _, val := range previous {
							fmt.Println(val.Name)
						}
					}
				}
			},
		},
		"nmap": {
			name:        "nmap",
			description: "Displays the next location",
			callback: func() {
				resultsCopy := make([]struct {
					Name string `json:"name"`
					URL  string `json:"url"`
				}, len(current.Results))
				copy(resultsCopy, current.Results)
				url, err := current.NextMap()
				if err != nil {
					fmt.Println(err)
				}
				if url != nil {
					locationCache.LocationAdd(*url, resultsCopy)
					for _, val := range current.Results {
						fmt.Println(val.Name)
					}
				}

			},
		},
	}
	scannedInput := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scannedInput.Scan()
		text := scannedInput.Text()
		clean := test.CleanInput(text)
		if clean == nil {
			fmt.Println("Unknown Command")
			continue
		}
		command, ok := cli[clean[0]]
		if ok == false {
			fmt.Println("Unknown Command")
			continue
		} else {
			switch command.name {
			case "exit":
				if len(clean) != 1 {
					fmt.Println("Unknown Command")
				} else {
					command.callback()
				}
			case "help":
				if len(clean) != 1 {
					fmt.Println("Unknown Command")
				} else {
					command.callback()
				}
			case "map":
				if len(clean) != 1 {
					fmt.Println("Unknown Command")
				} else {
					command.callback()
				}
			case "nmapb":
				if len(clean) != 1 {
					fmt.Println("Unknown Command")
				} else {
					command.callback()
				}
			case "nmap":
				if len(clean) != 1 {
					fmt.Println("Unknown Command")
				} else {
					command.callback()
				}
			case "explore":
				if len(clean) != 2 {
					fmt.Println("Require Two Arguments (Unsupported Digit and Single Characters)")
					continue
				} else {
					locationArea = clean[1]
					command.callback()
				}
			case "catch":
				if len(clean) != 2 {
					fmt.Println("Require Two Arguments (Unsupported Digit and Single Characters)")
					continue
				} else {
					pokemon = clean[1]
					command.callback()
				}
			}
		}
	}
}
