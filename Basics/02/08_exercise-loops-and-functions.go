package main

import (
	"fmt"
)

func Sqrt(x float64) float64 {
	z := 1.0
	i := 1
	for ; i < 10; {
		z -= (z*z - x) / (2 * z)
		fmt.Println(z)
		i ++
	}
	return z
}



func main() {
	Sqrt(10)
}