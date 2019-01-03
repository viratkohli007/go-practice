package main

import "fmt"

const pi = 3.14

type rect struct {
	height, width float64
}

func (r rect) area() float64 {

	return pi * r.height * r.width
}

func main() {

	r := rect{height: 10, width: 10}

	fmt.Println("area :", r.area())
}
