package main

import "fmt"

func DigitalRoot(n int) int {
	root := 0
	for n > 0 {
		root += n % 10
		n /= 10
	}

	if root > 9 {
		return DigitalRoot(root)
	}
	return root
}

func main() {

	fmt.Printf("DigitalRoot(20158672312301): %v\n", DigitalRoot(20158672312301))
}
