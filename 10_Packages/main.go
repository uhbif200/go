package main

import (
	"fmt"
	"lessons/10_Packages/my_math"
	"lessons/10_Packages/my_shit"
)

func main() {
	xs := []float64{1, 2, 3, 4}
	avg := my_math.Average(xs)
	fmt.Printf("avg: %v\n", avg)
	my_math.Hello()
	my_shit.Shit()
}
