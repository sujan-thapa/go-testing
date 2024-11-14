package main

import "fmt"

func Add(a, b int) int {
	return a + b
}

func Subtract(a, b int) int {
	return a - b
}

func Multiply(a, b int) int {
	return a * b
}

func main() {
	fmt.Println("Addition : ", Add(10, 20))
	fmt.Println("Subtraction : ", Subtract(10, 20))

	fmt.Println("Multiplication : ", Multiply(10, 20))

}
