package main

import (
	"fmt"
	"net/http"
)

func main() {
	/*http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello world")
	})
	http.HandleFunc("/cat", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello cat")
	})
	http.ListenAndServe(":9000",nil)*/
	myHandlerfunc := func(w http.ResponseWriter, r *http.Request){
		fmt.Fprintf(w,"Hello world")
		
	}
	
	http.HandleFunc("/",myHandlerfunc)
	http.HandleFunc("/cat",Dinesh)
	http.ListenAndServe(":9000",nil)
}
func Dinesh (w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w,"Hello cat")
}
