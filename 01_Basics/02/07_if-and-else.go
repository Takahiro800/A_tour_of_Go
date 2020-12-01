package main

import (
	"fmt"
	"math"
)

func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, lim)
	}
	return lim
}

func main() {
	//fmt.Prinln は改行しても1行出力になる
	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 20),
	)
}