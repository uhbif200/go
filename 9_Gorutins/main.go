package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

func f(n int) {
	for i := 0; i < n; i++ {
		fmt.Printf("i: %v\n", i)
	}
}

func suppler(c chan string) {
	for i := 0; ; i++ {
		c <- "Supply " + strconv.Itoa(i)
		sleep_time := time.Millisecond * time.Duration(rand.Intn(2000)+1000)
		time.Sleep(sleep_time)
	}
}

func printer(c chan string) {
	for {
		msg := <-c
		fmt.Printf("msg: %v\n", msg)
		sleep_time := time.Millisecond * time.Duration(rand.Intn(2000)+1000)
		time.Sleep(sleep_time)
	}
}

func supplierCreator(name string) func(c chan string) {
	return func(c chan string) {
		for i := 0; ; i++ {
			c <- "Supply by " + name + " " + strconv.Itoa(i)
			sleep_time := time.Millisecond * time.Duration(rand.Intn(2000)+1000)
			time.Sleep(sleep_time)
		}
	}
}

func main() {
	c := make(chan string, 10)

	go suppler(c)

	sup1 := supplierCreator("first")
	sup2 := supplierCreator("second")

	go sup1(c)
	go sup2(c)

	go printer(c)
	go printer(c)
	printer(c)
}
