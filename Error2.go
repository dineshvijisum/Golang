package main

import (
	"fmt"
	"os"
)

func main() {
	f, err := os.Open("inesh.txt")
	if err != nil {
		if e, d := err.(*os.PathError); d {
			fmt.Println("Failed to open file at path", e.Path)
			return
		}
		fmt.Println("Generic error", err)
		return
	}
	fmt.Println(f.Name(), "opened successfully")
}
