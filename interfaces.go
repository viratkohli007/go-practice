package main 
import "fmt"
import "math"

type geometry interface{

	area() float64
	perim() float64
}

type circle struct{

	radius float64
} 

type  rect struct{
	height, width float64
}

func (r rect) area() float64 {
	
   return r.height*r.width
}

func (r rect) perim() float64{
	return 2 * r.height * 2 * r.width
}

func (c circle) area() float64{
	return math.Pi*c.radius* c.radius
}

func (c circle) perim() float64{
	return 2 * math.Pi * c.radius
}

func measure(g geometry) {
	
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
}

func main() {
	
	a := rect{height: 10, width: 10}
	b := circle{radius: 10}

	measure(a)
	measure(b)

}