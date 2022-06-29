package main

import "fmt"

func main() {
	x := 5
	fmt.Printf("x: %v\n", x)
	increment(&x)
	fmt.Printf("x: %v\n", x)

	y := new(int)
	*y = 5
	fmt.Printf("*y: %v\n", *y)
	increment(y)
	fmt.Printf("*y: %v\n", *y)

	a := 1
	b := 2
	println("a, b =", a, b)
	swap(&a, &b)
	println("swapped a, b =", a, b)
}

func increment(x *int) {
	*x++
}

func swap(x *int, y *int) {
	tmp := *x
	*x = *y
	*y = tmp
}
