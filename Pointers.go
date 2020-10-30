package main

import "fmt"

func main() {
	/*i := 0
	//var i int = 0
	//fmt.Println(&i)
	var ptr *int
	ptr = &i
	fmt.Println(ptr)*/
	
	var i int=0
	var ptr *int
	var pptr **int
	ptr=&i
	pptr=&ptr
	fmt.Println(*ptr)
	fmt.Println(**pptr)
}
