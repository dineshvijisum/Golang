package main
import (
	"fmt"
	//"time"
)
func main(){
	myChannel :=make(chan string)
	go func (){
		myChannel <-"Hello"
		myChannel <-"World"
		}()
	text :=<-myChannel
	fmt.Println(text)
	fmt.Println(<)
}