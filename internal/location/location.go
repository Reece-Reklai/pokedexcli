package location

import (
	"encoding/json"
	"errors"
	"net/http"
)

type Location struct {
	Count    int     `json:"count"`
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func (location *Location) Location() error {
	res, err := http.Get("https://pokeapi.co/api/v2/location-area/")
	if err != nil {
		err = errors.New("Failed get general location request")
	}
	err = json.NewDecoder(res.Body).Decode(&location)
	if err != nil {
		err = errors.New("Failed to unmarshal json")
	}
	return err
}

func (location *Location) Prev_map() error {
	res, err := http.Get(*location.Previous)
	if err != nil {
		err = errors.New("Failed get general location request")
	}
	err = json.NewDecoder(res.Body).Decode(&location)
	if err != nil {
		err = errors.New("Failed to unmarshal json")
	}
	return err
}

func (location *Location) Next_map() error {
	res, err := http.Get(*location.Next)
	if err != nil {
		err = errors.New("Failed get general location request")
	}
	err = json.NewDecoder(res.Body).Decode(&location)
	if err != nil {
		err = errors.New("Failed to unmarshal json")
	}
	return err
}
