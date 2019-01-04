package main

import (
	"fmt"
	"time"
)

func main() {

	go say("virat")
	go say("kohli")

	// time.Sleep(time.Second)
}

func say(greet string) {

	for i := 0; i < 3; i++ {

		fmt.Println(greet)
	}

}
