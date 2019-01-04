package main

import "fmt"

func main() {
	done := make(chan string)
	fmt.Println("main function")
	go hello(done)
	fmt.Println(<-done)
}

func hello(done chan string) {
	fmt.Println("Hello world go routine")
	done <- "virat kohli"
}
