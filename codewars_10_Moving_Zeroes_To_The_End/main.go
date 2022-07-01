package main

import "fmt"

func MoveZeros(arr []int) []int {
	zeroCount := 0
	ans := []int{}
	for _, num := range arr {
		if num == 0 {
			zeroCount += 1
			continue
		}
		ans = append(ans, num)
	}

	for i := 0; i < zeroCount; i++ {
		ans = append(ans, 0)
	}
	// TODO: Program me
	return ans
}

func main() {
	fmt.Printf("MoveZeros([]int{0, 1, 2, 3, 4, 0, 5, 6, 7, 0}): %v\n", MoveZeros([]int{0, 1, 2, 3, 4, 0, 5, 6, 7, 0}))
	fmt.Printf("MoveZeros([]int{}): %v\n", MoveZeros([]int{}))
}
