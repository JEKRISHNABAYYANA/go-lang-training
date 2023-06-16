package main

import "fmt"

func palindrome(s string) bool {
	length := len(s)

	for i := 0; i < length/2; i++ {
		if s[i] != s[length-1-i] {
			return false
		}
	}

	return true
}

func main() {
	word1 := "aaaaaaaaaa"
	word2 := "bbbbbbbb"

	fmt.Println(word1, "is a palindrome:", palindrome(word1))
	fmt.Println(word2, "is a palindrome:", palindrome(word2))
}
