package main
import "fmt"

type person struct{

    name string
	age int
	}

func main() {
	
	fmt.Println(person{"bob", 20})

	fmt.Println(person{name: "Mayank", age:20})

	fmt.Println(person{name: "RIshi"})

	fmt.Println(&person{name: "prem", age: 21})

	s := person{
		name: "sean",
		age : 50}
	fmt.Println(s.name)

	sp := &s
	fmt.Println(sp.age)
}