package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z := 1.0
	i := 0.0

	for z - i > 0.0000001 || i - z > 0.0000001 {
		i = z
		z -= (z*z -x) / (2*z)
	}
	return z
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(math.Sqrt(2))
}