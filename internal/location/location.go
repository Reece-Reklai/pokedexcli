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

func (location *Location) Location() (err error) {
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

func (location *Location) PrevMap() (url *string, err error) {
	if location.Previous == nil {
		return nil, errors.New("Previous URL is currently unavailable")
	}
	res, err := http.Get(*location.Previous)
	if err != nil {
		err = errors.New("Failed get previous location request")
	}
	err = json.NewDecoder(res.Body).Decode(&location)
	if err != nil {
		err = errors.New("Failed to unmarshal json")
	}
	url = location.Previous
	return url, err
}

func (location *Location) NextMap() (url *string, err error) {
	if location.Next == nil {
		return nil, errors.New("Next URL is currently unavailable")
	}
	res, err := http.Get(*location.Next)
	if err != nil {
		err = errors.New("Failed get next location request")
	}
	err = json.NewDecoder(res.Body).Decode(&location)
	if err != nil {
		err = errors.New("Failed to unmarshal json")
	}
	url = location.Previous
	return url, err
}
