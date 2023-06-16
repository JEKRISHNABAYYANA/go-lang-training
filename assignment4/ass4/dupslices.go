package main

import "fmt"

func Duplicates(slice []int) []int {
	duplicates := make([]int, 0)
	m1 := make(map[int]bool)

	// Iterate over the slice and find duplicates
	for i := 0; i < len(slice); i++ {
		if m1[slice[i]] {
			continue
		}
		for j := i + 1; j < len(slice); j++ {
			if slice[i] == slice[j] {
				duplicates = append(duplicates, slice[i])
				m1[slice[i]] = true
				break
			}
		}
	}

	return duplicates
}

func main() {
	slice := []int{1, 2, 3, 4, 2, 3, 5, 6, 2, 2, 3, 4, 5, 2, 3}
	duplicates := Duplicates(slice)
	fmt.Println("Duplicates:", duplicates)
}
