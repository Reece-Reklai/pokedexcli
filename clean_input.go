package main

import (
	"strings"
)

func CleanInput(s string) []string {
	convert := []rune(s)
	var slice []string
	var tmp []rune
	for index, value := range convert {
		if index == 0 {
			if value != 32 {
				tmp = append(tmp, value)
			}
			continue
		}
		if value == 32 {
			if tmp != nil {
				conversion := string(tmp)
				slice = append(slice, strings.ToLower(conversion))
				tmp = nil
			}
			continue
		}
		tmp = append(tmp, value)
	}
	return slice
}
