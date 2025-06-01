package main

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

func (location *Location) location() error {
	res, err := http.Get("https://pokeapi.co/api/v2/location/")
	if err != nil {
		err = errors.New("Failed get general location request")
	}
	err = json.NewDecoder(res.Body).Decode(&location)
	if err != nil {
		err = errors.New("Failed to unmarshal json")
	}
	return err
}

func (location *Location) prev() error {
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

func (location *Location) next() error {
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
