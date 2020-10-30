package main

import "log"

type dii struct {
	Lat,d,k,l float64
}

var m map[string]dii

func main() {
	m = make(map[string]dii)
	m[""] = dii{
		40.68433, -74.39967,8,6.787555,
	}
	log.Println(m[""])
}