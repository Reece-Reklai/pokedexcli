package main

import "fmt"

func main() {
	hello := "hello world"
	convert := []rune(hello)
	var slice []string
	var tmp []rune
	for index, value := range convert {
		fmt.Printf("current value: %c\n", value)
		if value == 32 {
			conversion := string(tmp)
			slice = append(slice, conversion)
			tmp = nil
			continue
		}
		tmp = append(tmp, value)
		if index == len(hello)-1 {
			conversion := string(tmp)
			slice = append(slice, conversion)
		}
	}
	fmt.Println(slice)
	return
}
