package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type product struct {
	ID          int    `json:"id"`
	ProductName string `json:"productName"`
	Description string `json:"description"`
}

var products = []product{
	{ID: 1, ProductName: "Laptop", Description: "A laptop description"},
	{ID: 2, ProductName: "Smartphone", Description: "A Smartphone description"},
	{ID: 3, ProductName: "T-Shirt", Description: "A T-Shirt description"},
}

func main() {
	// GET Routes
	http.HandleFunc("/", rootHandler)
	http.HandleFunc("/products", productsHandler)

	// Listening to server
	fmt.Println("Server is listening on: 8080")
	err := http.ListenAndServe(":8080", nil)
	handleNil(err, "Error starting server: ")
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World!")
}

func productsHandler(w http.ResponseWriter, r *http.Request) {
	// Set the content type to JSON
	w.Header().Set("Content-Type", "application/json")

	// Convert the products slice to JSON
	productsJSON, err := json.Marshal(products)
	handleNil(err, "Error converting products to JSON: ")

	// Write the JSON response to the client
	w.Write(productsJSON)
}

func handleNil(err error, errMsg string) {
	if err != nil {
		fmt.Println(errMsg)
		panic(err)
	}
}
