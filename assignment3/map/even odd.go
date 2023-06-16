package map

import "fmt"

func findEvenOddElements(numbers []int) ([]int, []int) {
	even := []int{}
	odd := []int{}

	for _, num := range numbers {
		if num%2 == 0 {
			even = append(even, num)
		} else {
			odd = append(odd, num)
		}
	}

	return even, odd
}

func main() {
	// Example usage
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	evenElements, oddElements := findEvenOddElements(slice)

	fmt.Println("Even elements:", evenElements)
	fmt.Println("Odd elements:", oddElements)
}
