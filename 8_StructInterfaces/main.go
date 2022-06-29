package main

import (
	"fmt"
	"math"
	"strings"
)

type Circle struct {
	x, y, r float64
}

func (c *Circle) area() float64 {
	return math.Pi * math.Pow(c.r, 2)
}

type Rectangle struct {
	x, y, w, h float64
}

func (rect *Rectangle) area() float64 {
	return rect.h * rect.w
}

type Talking interface {
	speech() string
}

type Human struct {
	name string
}

func (*Human) speech() string {
	return "Ebal communistov!"
}

type Communist struct {
	Human
	was_gulaged bool
}

func (*Communist) speech() string {
	return "Ebal logiku!"
}

func showSpeeches(talkers ...Talking) {
	for _, t := range talkers {
		println(t.speech())
	}
}

type Bazar struct {
	talkers []Talking
}

func (b *Bazar) getPhrases() (ans []string) {
	for _, t := range b.talkers {
		ans = append(ans, t.speech())
	}
	return
}

func main() {
	c := Circle{x: 0, y: 0, r: 5}
	fmt.Printf("c: %v\n", c)

	c.x = 1
	c.y = 2
	fmt.Printf("c: %v\n", c)

	swap(&c.x, &c.y)
	fmt.Printf("c: %v\n", c)
	fmt.Printf("circleArea(c): %v\n", circleArea(c))
	fmt.Printf("area(c): %v\n", c.area())

	rect := Rectangle{x: 0, y: 0, w: 10, h: 10}
	fmt.Printf("area(rect): %v\n", rect.area())

	trozzkiy := Communist{Human: Human{name: "Lev"}, was_gulaged: true}

	durov := Human{name: "Pavel"}

	showSpeeches(&trozzkiy, &durov)

	b := Bazar{talkers: []Talking{&trozzkiy, &durov}}
	fmt.Printf("Dialog na bazare: %v", strings.Join(b.getPhrases(), " "))

}

func swap(x, y *float64) {
	tmp := *x
	*x = *y
	*y = tmp
}

func circleArea(c Circle) float64 {
	return math.Pi * math.Pow(c.r, 2)
}
