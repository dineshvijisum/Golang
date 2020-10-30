package main

import "fmt"

/*func main(){
x := max(5,"Dinesh")
fmt.Println(x)
}
func max(x int,y string)(int,string){
	/*if {
	return x
	}else{
		return y*/

func main() {
	x := 10
	y := 20
	x, y = max(x, y)
	fmt.Println(x, y)
}
func max(x, y int) (int, int) {
	return y, x
}
