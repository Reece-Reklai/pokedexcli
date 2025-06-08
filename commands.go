package main

import (
	"fmt"
	"os"
)

func commandExit() {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return
}
func commandHelp() {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println("")
	fmt.Println("help: Displays a help message")
	fmt.Println("map: Locations that can be visited within the games. Locations make up sizable portions of regions, like cities or routes.")
	fmt.Println("nmap: The Next Locations")
	fmt.Println("nmapb: The Prev Locations")
	fmt.Println("explore: Pokemon encountered in location-area")
	fmt.Println("catch: Throw a Pokeball at a Pokemon")
	fmt.Println("inspect: Inspect Pokemon from Pokedex")
	fmt.Println("pokedex: What is currently inside your Pokedex")
	fmt.Println("exit: Exit the Pokedex")
	return
}
