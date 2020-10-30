package main

import ("fmt"
       "reflect")
func divide(x, y int64) int64{
	var res = x / y
	return res
}
func main() {
	var n1 int64 = 5
	var n2 int64 = 3
	var a float64 = float64(n1)
	var b float64 = float64(n2)
	var n float64=a/b
	fmt.Println("\n", reflect.TypeOf(a))
	fmt.Println("\n", reflect.TypeOf(b))
	fmt.Println("\n", reflect.TypeOf(n1))
   fmt.Println(n)
}
