package main

import (
	"fmt"
	"strings"
)

func main() {
	const hell string = "Hello"
	fmt.Println(hell)
	var x string = "Lol"
	var y = x
	z := y

	fmt.Println(x, y, z)

	k := strings.Split(x, "o")

	fmt.Printf("k: %v\n", k)

	fmt.Println("Zero farengeit is", farToCels(0), "celsium")
}

func farToCels(deg int) int {
	return (deg - 32) * 5 / 9
}
