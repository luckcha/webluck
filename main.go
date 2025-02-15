package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Product struct {
	Id       int
	Name     string
	Quantity int
	Price    float64
}

var Products []Product

func homepage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hey there")
	log.Println("Endpoint hit:homepage")
}
func returnallproduct(w http.ResponseWriter, r *http.Request) {
	log.Println("Endpoint hit:returnallproduct")
	json.NewEncoder(w).Encode(Products)
}
func getproduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]

	id, err := strconv.Atoi(key)
	if err != nil {
		http.Error(w, "Invalid product ID", http.StatusBadRequest)
		return
	}

	for _, product := range Products {
		if (product.Id) == id {
			json.NewEncoder(w).Encode(product)

		}
	}
}

func handleRequest() {
	Myrouter := mux.NewRouter().StrictSlash(true)
	Myrouter.HandleFunc("/", homepage)
	Myrouter.HandleFunc("/product/{id}", getproduct)
	Myrouter.HandleFunc("/products", returnallproduct)
	http.ListenAndServe("localhost:5050", Myrouter)

}

func main() {
	Products = []Product{
		{Id: 1, Name: "Laptop", Quantity: 80, Price: 32000.50},
		{Id: 2, Name: "Refrigerator", Quantity: 70, Price: 84000.70},
		{Id: 3, Name: "TV", Quantity: 95, Price: 23000.50},
		{Id: 4, Name: "Mobile", Quantity: 68, Price: 64200.60},
	}
	handleRequest()
}
