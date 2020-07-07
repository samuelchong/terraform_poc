package main

import (
	"fmt"
)

//func helloWorld(w http.ResponseWriter, r *http.Request) {
func helloWorld() {
	fmt.Println("Hello World!123")
}

func main() {
	// http.HandleFunc("/", helloWorld)
	// http.ListenAndServe(":8080", nil)
	helloWorld()
}
