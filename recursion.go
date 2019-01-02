package main
import "fmt"

func main(){

fmt.Println(recur(5))

}

func  recur(n int) int {
	if (n == 1){
		return n
	}else{
		return n*recur(n-1)
	}
}