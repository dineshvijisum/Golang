package main

import "fmt"

func main() {
	/*var a = []int{5, 6, 7, 8, 9}
	a[2] = 1
	var b = a[1]
	b = 10
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(a)*/
	var num = []int{5, 6, 7, 8, 9}
	for i := 0; i < 5; i++ {
		fmt.Printf("Value at %d is %d\n",num[i],i)
	}

	/*Range Function
var num =[]int{3,4,5,6,7}
for i,x :=range num{
	fmt.Printf("valu at %d is %d\n",i,x)
}*/
}

