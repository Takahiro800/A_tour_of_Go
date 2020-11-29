package main

import (
	"fmt"
	"math"
)

type vertex2 struct {
	X, Y float64
}

func Abs(v vertex2) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := vertex2{3, 4}
	fmt.Println(Abs(v))
}
