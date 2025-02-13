package main

import (
	"fmt"
	"log"
	"net/http"
)

var anmo struct {
	Id          int
	Name        string
	Address     string
	Designation string
	Salary      float64
}

func homepage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hey there")
	log.Println("Endpoint hit:homepage")
}

func main() {
	http.HandleFunc("/", homepage)
	http.ListenAndServe("localhost:5050", nil)
}
