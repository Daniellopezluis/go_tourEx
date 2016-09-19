package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	zn := x
	var auxzn float64
	i := 0
	const TOLERANCE = 0.000000000000001
	for math.Abs(zn-auxzn) > TOLERANCE {
		auxzn = zn
		zn -= (zn*zn - x) / (2 * zn)
		i++
	}
	fmt.Printf("Number of loops executed: \t%d\n", i)
	return zn
}

func main() {
	fmt.Printf("My sqrt function: \t\t%g\n", Sqrt(2))
	fmt.Printf("Math lib's sqrt function: \t%g\n", math.Sqrt(2))
}
