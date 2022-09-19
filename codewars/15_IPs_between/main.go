package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func getBytes(ip string) []int {
	splitted := strings.Split(ip, ".")
	var numbers []int
	for _, num := range splitted {
		number, _ := strconv.Atoi(num)
		numbers = append(numbers, number)
	}
	return numbers
}

func IpsBetween(start, end string) int {
	bytes_start := getBytes(start)
	bytes_end := getBytes(end)
	between := 0
	for i, _ := range bytes_start {
		between += (bytes_end[i] - bytes_start[i]) * int(math.Pow(256, float64(3-i)))
	}
	return (int(math.Abs(float64(between))) - 1)
}

func main() {
	fmt.Printf("IpsBetween(\"192.168.1.1\", \"192.168.1.2\"): %v\n", IpsBetween("192.168.2.1", "192.168.1.1"))
}
