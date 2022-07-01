package main

import (
	"fmt"
	"regexp"
	"sort"
)

func InArray(array1 []string, array2 []string) []string {
	// your code

	allKeys := make(map[string]bool)
	list := []string{}
	for _, item := range array1 {
		if _, value := allKeys[item]; !value {
			allKeys[item] = true
			list = append(list, item)
		}
	}

	ans := []string{}
	for _, word := range list {
		for _, test := range array2 {
			if matched, _ := regexp.MatchString(word, test); matched {
				ans = append(ans, word)
				break
			}
		}
	}

	sort.Strings(ans)

	return ans
}

func main() {
	a1 := []string{"arp", "live", "strong"}
	a2 := []string{"lively", "alive", "harp", "sharp", "armstrong"}
	a3 := InArray(a1, a2)
	fmt.Printf("a3: %v\n", a3)
}
