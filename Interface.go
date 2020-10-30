package main
import "fmt"
import "math"

type shape interface{
	area() float64
}
type circle struct{
	radius float64
}
func(c circle)area() float64{
	return math.pi * c.radius * c.radius
}
type rectangle struct{
	height,width float64
}
func(r.rectangle)area() float64{
	return r.width*r.height
}
func getArea(s shape)float64{
	return s.area()
}
unc main(){
	c :=circle{radius:10}
	r :=rectangle{height:2,width:4}
	fmt.Println(getArea(c))
	fmt.Println(getArea(r))
}