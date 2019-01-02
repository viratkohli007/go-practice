package main
import "fmt"

func sum(nums... int){
    fmt.Println(nums, " ")
	total := 0
	for _,num := range nums{
	   
	   total += num
	}
	fmt.Print(total,"\n")
}

func main(){
	
   sum (1,23,3)
   sum (2,34,3)

   arr1 := []int {3,4,5,6,7}
   sum(arr1...)
}