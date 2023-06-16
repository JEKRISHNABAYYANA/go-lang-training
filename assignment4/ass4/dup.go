package main

import "fmt"

func findDuplicates(slice []int) []int {
	duplicates := make([]int, 0)

	// Iterate over the slice and find duplicates
	for i := 0; i < len(slice); i++ {
		for j := i + 1; j < len(slice); j++ {
			if slice[i] == slice[j] {
				duplicates = append(duplicates, slice[i])
			}
		}
	}

	return duplicates
}

func main() {
	slice := []int{1, 2, 3, 4, 1, 2, 3, 4, 1, 2, 3, 4, 1, 2, 3, 4}
	duplicates := findDuplicates(slice)
	fmt.Println("Duplicates:", duplicates)
}
