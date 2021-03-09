package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	e := 0.00000000000000025
	zold, z :=0.0, 1.0
	for z-zold >= e || z-zold <= -e {
		zold = z
		z -= (z*z - x) / (2*z)
	}
	return z
}

func main() {
	fmt.Println(Sqrt(2), math.Sqrt(2))
}
