package main
import "fmt"

func val()(int, int) {
	return 4,5
}

func main() {
	
	a, b := val()
	fmt.Println(a)
	fmt.Println(b)

	_, c := val()
	fmt.Println(c)
	//fmt.Println(_)   //error
}