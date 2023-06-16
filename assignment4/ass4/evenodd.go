package main

import "fmt"

func main() {
	var x = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	var even []int
	var odd []int

	for i := range x {
		if x[i]%2 == 0 {
			even = append(even, x[i])
		} else {
			odd = append(odd, x[i])
		}
	}

	fmt.Println("Even Numbers:", even)
	fmt.Println("Odd Numbers:", odd)
}
