package main
import "fmt"

func sum(a int, b int)int{
   return a+b
}
func sum3(a,b,c int) int{
	 return a+b+c
}
func main() {
	
	res := sum(4,5)
fmt.Println("res is: ", res)
res = sum3(3,4,5)
fmt.Println("now res: ",res)
}