package main

import "fmt"

func main() {
	a := 5
	b := 2

	fmt.Println("Before swap:", a, b)

	a, b = b, a

	//a = b - a
	//b = b - a
	//a = a + b

	fmt.Println("After swap:", a, b)
}
