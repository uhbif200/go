package main

import (
	"fmt"
	"os"
)

func main() {
	evenGenerator := makeEvenGenerator()
	fmt.Printf("evenGenerator(): %v\n", evenGenerator())
	fmt.Printf("evenGenerator(): %v\n", evenGenerator())
	fmt.Printf("evenGenerator(): %v\n", evenGenerator())

	fmt.Printf("factorial(10): %v\n", factorial(10))
	fmt.Printf("factorialRecursive(10): %v\n", factorialRecursive(10))
	writeToFile()

	x := []int{2, 4, 10}
	fmt.Printf("sum(x): %v\n", sum(x))
	fmt.Println(half(4))
	fmt.Println(half(7))
	fmt.Printf("fib(10): %v\n", fib(15))

	fibfib := fibonaccinator()
	fmt.Printf("fibfib(10): %v\n", fibfib(15))
}

func makeEvenGenerator() func() int {
	i := int(0)
	return func() (ret int) {
		ret = i
		i += 2
		return
	}
}

func factorial(x int) int {
	ret := x
	for i := x - 1; i > 0; i-- {
		ret *= i
	}
	return ret
}

func factorialRecursive(x int) int {
	if x <= 1 {
		return 1
	}
	return x * factorialRecursive(x-1)
}

func writeToFile() {
	f, err := os.OpenFile("zalupa.txt", os.O_RDWR, 0644)
	if err != nil {
		println("cannot open file, trying to create")
		os.Create("zalupa.txt")
	}
	f, err = os.OpenFile("zalupa.txt", os.O_RDWR, 0644)
	if err != nil {
		println("cannot create file")
		return
	}

	f.WriteString("zalupa byl zdes")
	defer func() {
		f.Close()
	}()
}

func sum(x []int) int {
	var ret int
	for i := 0; i < len(x); i++ {
		ret += x[i]
	}
	return ret
}

func half(x int) (x_half int, even bool) {
	x_half = x / 2
	if x%2 == 0 {
		even = true
	}
	return
}

func fib(x int) int {
	if x == 0 {
		return 0
	}
	if x == 1 {
		return 1
	}
	return fib(x-1) + fib(x-2)
}

func fibonaccinator() func(x int) uint64 {
	buffer := make([]uint64, 100)

	var retFunc func(x int) uint64

	retFunc = func(x int) uint64 {
		if x == 0 {
			return 0
		}
		if x == 1 {
			return 1
		}
		pos := x - 1
		if buffer[pos] != 0 {
			return buffer[pos]
		}

		buffer[pos] = retFunc(x-2) + retFunc(x-1)
		return buffer[pos]
	}
	return retFunc
}
