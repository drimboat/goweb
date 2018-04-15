package main

import(
	"fmt"
	"net/http"
)

func main(){
	fmt.Print("Started")
	http.HandleFunc("/", indexHandler)
	http.ListenAndServe(":8080", nil)
}

func indexHandler(w http.ResponseWriter, r *http.Request){
	fmt.Fprintln(w,"Hello world")
}