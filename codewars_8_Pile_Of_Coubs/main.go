package main

import (
	"fmt"
)

func Pow(a int, b int) int {
	ans := a
	for i := 2; i <= b; i++ {
		ans *= a
	}
	return ans
}

func FindNb(m int) int {
	curVol := m
	curCube := 1
	for {
		if curVol == 0 {
			return curCube - 1
		}
		if curVol < 0 {
			return -1
		}
		curVol -= Pow(curCube, 3)
		curCube += 1
	}
}

func main() {
	fmt.Printf("FindNb(1): %v\n", FindNb(9))
}
