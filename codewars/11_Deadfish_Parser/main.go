package main

import (
	"fmt"
)

func Parse(data string) []int {
	//TODO: write your code here
	ans := []int{}
	i := 0
	for _, symbol := range data {
		switch symbol {
		case 'i':
			i++
		case 's':
			i *= i
		case 'd':
			i--
		case 'o':
			ans = append(ans, i)
		}
	}
	return ans
}

func main() {
	fmt.Printf("Parse(\"iiisdoso\"): %v\n", Parse("iiisdoso"))
}
