package main

import "fmt"

func main() {

	a := []int{2, 3, 4}
	sum := 0
	for _, num := range a {
		sum += num
	}
	fmt.Println("sum: ", sum)

	for i, num := range a {
		if num == 3 {
			fmt.Println("index: ", i)
		}
	}
	m := map[string]string{"a": "apple", "c": "cat"}
	for m1, m2 := range m {
		fmt.Println("key: ", m1, "value: ", m2)
	}

	for i, c := range "go" {
		fmt.Println(i, c)
	}

}
