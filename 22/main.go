package main

import (
	"fmt"
	"math/big"
)

func main() {
	a := big.NewFloat(10_000_000_000_000.500)
	b := big.NewFloat(5_000_000_000_000.250)

	mul := big.NewFloat(0)
	mul.Mul(a, b)
	fmt.Printf("a * b = %f\n", mul)

	div := big.NewFloat(0)
	div.Quo(a, b)
	fmt.Printf("a / b = %f\n", div)

	sum := big.NewFloat(0)
	sum.Add(a, b)
	fmt.Printf("a + b = %f\n", sum)

	diff := big.NewFloat(0)
	diff.Sub(a, b)
	fmt.Printf("a - b = %f\n", diff)
}
