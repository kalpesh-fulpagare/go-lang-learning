package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) (z float64) {
	fmt.Println("Input: ", x)
	z = 1.0
	var y float64
	for i := 0; i < 10; i++ {
		z -= (z*z - x) / (2*z)
		y = z * z
		if x == y {
			return
		} else if y > x {
			z = z - 0.10
		} else {
			z = z + 0.10
		}
		fmt.Printf("\nIteration: %d Value of z: %g Value of y: %g", i, z, y)
	}
	return
}

func main() {
	fmt.Println("\n\nSquare Root: ", Sqrt(2.0))
	fmt.Println("Square Root: ", math.Sqrt(2))
}
