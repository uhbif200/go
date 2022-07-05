package main

import (
	"fmt"
	"regexp"
)

func Solution(str string) []string {
	var ans []string
	re := regexp.MustCompile(".{2}")
	ans = re.FindAllString(str, -1)
	if len(str)%2 != 0 {
		ans = append(ans, string(str[len(str)-1])+"_")
	}
	return ans
}

func main() {
	strs := Solution("asdffsdfasdff")
	fmt.Printf("strs: %v\n", strs)
}
