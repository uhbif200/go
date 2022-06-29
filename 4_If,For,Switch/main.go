package main

import "fmt"

func main() {
	print10Nums()
	print10NumsEvenOdd()
	print10NumsFull()
	fizzBuzz()
	// var x int
	// fmt.Scan(x)
}

func print10Nums() {
	fmt.Println("Print 10 nums")
	for i := 0; i < 10; i++ {
		fmt.Printf("i: %v\n", i)
	}
}

func print10NumsEvenOdd() {
	fmt.Println("Print 10 nums even odd")
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			fmt.Println("Even")
		} else {
			fmt.Println("Odd")
		}
	}
}

func print10NumsFull() {
	fmt.Println("Print 10 nums full names")
	for i := 0; i < 10; i++ {
		switch i {
		case 0:
			fmt.Println("Zero")
		case 1:
			fmt.Println("One")
		case 2:
			fmt.Println("Two")
		case 3:
			fmt.Println("Three")
		case 4:
			fmt.Println("Four")
		case 5:
			fmt.Println("Five")
		default:
			fmt.Println("Unknown Number")
		}
	}
}

func fizzBuzz() {
	fmt.Println("Print fizzbuzz")
	for i := 0; i < 100; i++ {
		var printed = false
		if i%3 == 0 {
			fmt.Print("Fizz")
			printed = true
		}
		if i%5 == 0 {
			fmt.Print("Buzz")
			printed = true
		}
		if printed {
			fmt.Print("\n")
		}
	}
}
