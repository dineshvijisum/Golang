package main

import (  
    "fmt"
    "path/filepath"
)

func main() {  
    files, _ := filepath.Glob("4")
    fmt.Println("matched files", files)
}