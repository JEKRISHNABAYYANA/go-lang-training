// Given two binary strings a and b, return their sum as a binary string.

// Example 1:

// Input: a = "11", b = "1"
// Output: "100"
// Example 2:

// Input: a = "1010", b = "1011"
// Output: "10101"

// Constraints:

// 1 <= a.length, b.length <= 104
// a and b consist only of '0' or '1' characters.
// Each string does not contain leading zeros except for the zero itself.

package main

import (
	"fmt"
)

func addBinary(a string, b string) string {
	result := ""
	carry := 0
	i, j := len(a)-1, len(b)-1

	for i >= 0 || j >= 0 || carry > 0 {
		sum := carry

		if i >= 0 {
			sum += int(a[i] - '0')
			i--
		}
		if j >= 0 {
			sum += int(b[j] - '0')
			j--
		}

		result = fmt.Sprintf("%d%s", sum%2, result)
		carry = sum / 2
	}

	return result
}

func main() {
	a := "11"
	b := "1"
	fmt.Println(addBinary(a, b)) // Output: 100

	a = "1010"
	b = "1011"
	fmt.Println(addBinary(a, b)) // Output: 10101
}
