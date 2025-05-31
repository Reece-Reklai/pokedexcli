package main

import (
	"strings"
)

func CleanInput(s string) []string {
	convert := []rune(s)
	var slice []string
	var tmp []rune
	if s == "" {
		return nil
	}
	if len(s) == 1 {
		slice = append(slice, strings.ToLower(s))
		return slice
	}
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
		if index == len(s)-1 {
			if tmp != nil {
				tmp = append(tmp, value)
				conversion := string(tmp)
				slice = append(slice, strings.ToLower(conversion))
				tmp = nil
			}
		}
		tmp = append(tmp, value)
	}
	return slice
}
