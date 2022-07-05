package main

import "fmt"

func FindOutlier(integers []int) int {
	var (
		odd  []int
		even []int
	)

	for _, number := range integers {
		if number%2 == 0 {
			even = append(even, number)
		} else {
			odd = append(odd, number)
		}

		if len(odd) > 1 && len(even) > 0 {
			return even[0]
		}
		if len(even) > 1 && len(odd) > 0 {
			return odd[0]
		}
	}
	return 0
}

func main() {
	fmt.Printf("FindOutlier([]int{2, 4, 5, 6, 8}): %v\n", FindOutlier([]int{2, 4, 5, 6, 8}))
	fmt.Printf("FindOutlier([]int{1, 3, 5, 6, 3}): %v\n", FindOutlier([]int{1, 3, 5, 6, 3}))

}
