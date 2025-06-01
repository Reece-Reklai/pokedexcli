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
