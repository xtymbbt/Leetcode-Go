package main

import (
	"fmt"
	"math"
)

func climbStairs(n int) int {
	sqrt5 := math.Sqrt(5)
	pow1 := math.Pow((1+sqrt5)/2, float64(n+1))
	pow2 := math.Pow((1-sqrt5)/2, float64(n+1))
	return int(math.Round((pow1 - pow2) / sqrt5))
}

func main() {
	fmt.Println(climbStairs(3))
}