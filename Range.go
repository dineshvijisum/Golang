package main
import "fmt"

func main() {

	Dinesh := []string{"Go", "C", "C++", "Java"}
	fmt.Println("W3Adda - Go For Each Loop.")
	for i,s := range Dinesh{
		fmt.Println(i, s)
	}
}