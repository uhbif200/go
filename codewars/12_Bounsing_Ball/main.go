package main

import (
	"fmt"
)

func BouncingBall(h, bounce, window float64) int {
	if !(h > 0 && bounce > 0 && bounce < 1 && window < h) {
		return -1
	}
	count := 1

	for {
		h *= bounce
		if h > window {
			count += 2
		} else {
			return count
		}
	}
}

func main() {
	fmt.Printf("BouncingBall(3, 1, 1.5): %v\n", BouncingBall(3, 0.66, 1.5))
}
