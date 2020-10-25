package main

import (
	"fmt"
	"math"
)

func main() {
	var a, v0, s0, t float64
	fmt.Printf("Enter values of a: ")
	fmt.Scan(&a)

	fmt.Printf("Enter values of v0: ")
	fmt.Scan(&v0)

	fmt.Printf("Enter values of s0: ")
	fmt.Scan(&s0)

	fmt.Printf("Enter values of t: ")
	fmt.Scan(&t)

	s := genDisplaceFunc(a, v0, s0)

	fmt.Printf("The displacemnt after %vs is: %v", t, s(t))
}

func genDisplaceFunc(a float64, v0 float64, s0 float64) func(float64) float64 {
	fn := func(t float64) float64 {
		return (0.5 * a * math.Pow(t, 2)) + (v0 * t) + s0
	}
	return fn
}
