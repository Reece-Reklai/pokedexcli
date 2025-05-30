package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	for {
		fmt.Print("Pokedex > ")
		scanned_input := bufio.NewScanner(os.Stdin)
		scanned_input.Scan()
		text := scanned_input.Text()
		clean := CleanInput(text)
		fmt.Printf("Your command was: %v\n", clean[0])
	}
}
