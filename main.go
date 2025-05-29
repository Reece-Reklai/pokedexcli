package main

import (
	"fmt"
	"strings"
)

func main() {
	hello := "hello woRld COMputER"
	convert := []rune(hello)
	var slice []string
	var tmp []rune
	for index, value := range convert {
		fmt.Printf("current value: %c\n", value)
		if value == 32 {
			conversion := string(tmp)
			slice = append(slice, strings.ToLower(conversion))
			tmp = nil
			continue
		}
		tmp = append(tmp, value)
		if index == len(hello)-1 {
			conversion := string(tmp)
			slice = append(slice, strings.ToLower(conversion))
		}
	}
	fmt.Println(slice)
	return
}
