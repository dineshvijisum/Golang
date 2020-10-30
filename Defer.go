package main

import (
	"fmt"
)

func main() {
	defer printText("Dinesh")
	printText("hi")
	defer printText("kumar")
	defer printText("vs")
}
func printText(p string) {
	fmt.Println(p)
}
