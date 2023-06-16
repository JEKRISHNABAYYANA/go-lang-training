package main

import (
	"fmt"
	"strings"
	"unicode"
)

func MiddleIndex(str string, char rune) int {
	count := strings.Count(strings.ToLower(str), string(unicode.ToLower(char)))
	if count == 0 {
		return -1 // Character not found in the string
	}

	occurrence := 0
	for i, c := range str {
		if unicode.ToLower(c) == unicode.ToLower(char) {
			occurrence++
			if occurrence == (count+1)/2 {
				return i
			}
		}
	}

	return (count - 1) / 2 // Return the floor value when the count is odd
}

func main() {
	str := "this is a sample string"
	char := 'S'

	index := MiddleIndex(str, char)
	if index == -1 {
		fmt.Println("Character not found in the string.")
	} else {
		fmt.Printf("The middle index of character '%c' is %d.\n", char, index)
	}
}
