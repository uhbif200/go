package main

import (
	"fmt"
	"regexp"
	"strings"
)

//TODO packet unicode can check is symbol letter

func wave(words string) []string {
	var vibe []string
	// Your code here and happy coding!
	for index, letter := range words {
		rg := regexp.MustCompile("[a-zA-Z]")
		if rg.MatchString(string(letter)) {
			if index == 0 {
				vibe = append(vibe, strings.ToUpper(string(letter))+words[index+1:])
			} else {
				vibe = append(vibe, words[0:index]+strings.ToUpper(string(letter))+words[index+1:])
			}
		}
	}
	if vibe == nil {
		return make([]string, 0)
	}

	return vibe
}

func main() {
	fmt.Printf("wave(\"holla amigos\"): %v\n", wave("hs"))
}
