package main

import (
	"fmt"
	"net/http"
)
func helloHandler( w http.ResponseWriter, r *http.Request){
	fmt.Fprintln(w,"hello world")
}
func main() {
	mux:= http.NewServeMux() // router
	mux.HandleFunc("/hello",helloHandler)
	fmt.Println("server is running on :3000") //route
	
	err := http.ListenAndServe(":3000",mux) // expose port 

	if(err != nil){
		fmt.Println("error",err) // error 
	}
	
}