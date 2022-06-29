package main

import (
	"fmt"
	"math"
)

type Circle struct {
	x, y, r float64
}

func main() {
	c := Circle{x: 0, y: 0, r: 5}
	fmt.Printf("c: %v\n", c)

	c.x = 1
	c.y = 2
	fmt.Printf("c: %v\n", c)

	swap(&c.x, &c.y)
	fmt.Printf("c: %v\n", c)
	fmt.Printf("circleArea(c): %v\n", circleArea(c))

	z := fun.fibonaccinator()
	z(10)
}

func swap(x, y *float64) {
	tmp := *x
	*x = *y
	*y = tmp
}

func circleArea(c Circle) float64 {
	return math.Pi * math.Pow(c.r, 2)
}
