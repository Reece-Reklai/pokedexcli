package main

import (
	"fmt"
)

func commandHelp() {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println("")
	fmt.Println("help: Displays a help message")
	fmt.Println("map: Locations that can be visited within the games. Locations make up sizable portions of regions, like cities or routes.")
	fmt.Println("exit: Exit the Pokedex")
	return
}
