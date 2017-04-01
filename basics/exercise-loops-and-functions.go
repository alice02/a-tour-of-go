package main

import (
	"fmt"
	"math"
)

const eps = 1e-9

func Sqrt(x float64) float64 {
	z, z_prev := 1.0, 0.0
	for {
		z = z - (z*z-x)/(2*z)
		if math.Abs(z_prev-z) < eps {
			break
		}
		z_prev = z
	}
	return z
}

func main() {
	fmt.Println("Newton method :", Sqrt(2))
	fmt.Println("math.Sqrt     :", math.Sqrt(2))
}
