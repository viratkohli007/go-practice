package main

import "fmt"

func main() {

	num := 10
	sq := make(chan int)
	cu := make(chan int)

	go calSqure(num, sq)
	go calCube(num, cu)

	squres, cubes := <-sq, <-cu
	fmt.Println("final output: ", squres, cubes)

}

func calSqure(num int, sqop chan int) {

	sum := 0

	sum += num * num

	sqop <- sum

}

func calCube(num int, cubeop chan int) {

	sum := 0

	sum += num * num * num

	cubeop <- sum

}
