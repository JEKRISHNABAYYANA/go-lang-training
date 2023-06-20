package main

import (
	"fmt"
	"go-lang-training/assignment3/math"
)

func main() {
	var x, y, z, a int
	x, y, z = math.Addition(50, 20)
	fmt.Println("sum of the two numbers is", "integer 1 =", x, "integer 2 =", y, "addition of both integers= ", z)

	x, y, z = math.Subtraction(70, 35)
	fmt.Println("subtraction  of the two numbers is", "integer 1 =", x, "integer 2 =", y, "subtraction of both integers= ", z)

	x, y, z = math.Multiplication(12, 42)
	fmt.Println("Multiplication of the two numbers is", "integer 1 =", x, "integer 2 =", y, "multiplication of both integers= ", z)

	x, y, z, a = math.Divison(20, 10)
	fmt.Println("Divison of the two numbers is", x, y, "quotient =", z, "remainder =", a)

}
