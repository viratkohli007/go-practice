package main
import "fmt"

// func sum(nums... int){
//     fmt.Println(nums, " ")
// 	total := 0
// 	for _,num := range nums{
	   
// 	   total += num
// 	}
// 	fmt.Print(total,"\n")
// }

func makeEvenGen() func() int{
	 i:= 1
	 return func() (ret int){
	 	 ret = i
	 	 i += 2
	 	 return 
	 }
}

func printInt() func() int {
	i := 1
	return func () int{
		i++
		return i
	}
}

func main(){
   
    nextInt := printInt()


    nextEven := makeEvenGen()
	add := func (a, b int) int {
		return a+b
	}

	fmt.Println(add(3,4))
      i :=1
	increment := func () int{
		i++
		return i
	}

	fmt.Println(increment())
	fmt.Println(nextEven())	
	fmt.Println(nextEven())	
	fmt.Println(nextInt())
     	fmt.Println(nextInt())

}


