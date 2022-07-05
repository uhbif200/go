package main

import (
	"fmt"
	"strings"
)

func format(num *int) {
	if *num < 0 {
		*num = 0
		return
	}
	if *num > 255 {
		*num = 255
		return
	}
}

func RGB(r, g, b int) string {
	// Your code here
	format(&r)
	format(&g)
	format(&b)

	ans := []string{}

	ans = append(ans, strings.ToUpper(fmt.Sprintf("%02x", r)))
	ans = append(ans, strings.ToUpper(fmt.Sprintf("%02x", g)))
	ans = append(ans, strings.ToUpper(fmt.Sprintf("%02x", b)))

	return strings.Join(ans, "")
}

func main() {
	fmt.Printf("RGB(0, 0, 0): %v\n", RGB(0, 0, 0))
	fmt.Printf("RGB(0, 0, 0): %v\n", RGB(255, 255, 255))
	fmt.Printf("RGB(148, 0, 211): %v\n", RGB(148, 0, 234234))

}
