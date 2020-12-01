package main

import (
	"fmt"
	"math"
)

//　乗算とlimの小さい方を返す
func lim_pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	}
	return lim
}

func main() {
	fmt.Println(
		lim_pow(3, 2, 10),
		lim_pow(3, 3, 20),
	)
}
