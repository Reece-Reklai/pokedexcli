package player

import (
	"encoding/json"
	"errors"
	"fmt"
	"math/rand"
	"net/http"
)

type Player struct {
	Pokedex map[string]Pokemon
}

func attempt(attempt Pokemon) bool {
	diffcultyFactor := rand.Intn(attempt.BaseExperience)
	if diffcultyFactor < (attempt.BaseExperience / 2) {
		return true
	}
	return false
}

func (caught *Player) Catch(name string) (err error) {
	var pokemon Pokemon
	res, err := http.Get("https://pokeapi.co/api/v2/pokemon/" + name)
	if err != nil {
		err = errors.New("Unable to recieve Pokemon Information")
	}
	err = json.NewDecoder(res.Body).Decode(&pokemon)
	if err != nil {
		err = errors.New("Pokemon does not exist")
	} else {
		_, ok := caught.Pokedex[name]
		if ok == true {
			fmt.Println("Pokemon already caught")
		} else {
			fmt.Printf("Throwing a Pokeball at %s...\n", name)
			verdict := attempt(pokemon)
			if verdict == true {
				fmt.Printf("%s was caught!\n", name)
				if caught.Pokedex == nil {
					caught.Pokedex = make(map[string]Pokemon)
					caught.Pokedex[name] = pokemon
				} else {
					caught.Pokedex[name] = pokemon
				}
			} else {
				fmt.Printf("%s escaped!\n", name)
			}
		}
	}
	return err
}
