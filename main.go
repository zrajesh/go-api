package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type product struct {
	ID          int    `json:"id"`
	ProductName string `json:"productName"`
	Description string `json:"description"`
}

// User struct to represent a user
type User struct {
	Username string `form:"username"`
	Email    string `form:"email"`
	Password string `form:"-"`
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

	// POST Routes
	http.HandleFunc("/create-product", createProductHandler)
	http.HandleFunc("/create-user", createUserHandler)

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

func createProductHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	// Read the request body
	body, err := io.ReadAll(r.Body)
	handleNil(err, "Error reading request body: ")

	// Unmarshal JSON data into a product struct
	var newProduct product
	err = json.Unmarshal(body, &newProduct)
	handleNil(err, "Error unmarshalling JSON: ")

	// Assign a unique ID to the new product (you might want to use a more robust ID generation mechanism)
	newProduct.ID = len(products) + 1

	// Add the new product to the products slice
	products = append(products, newProduct)

	// Respond with the created product
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newProduct)
}

func createUserHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	// Parse the form data
	err := r.ParseForm()
	handleNil(err, "Error parsing form data: ")

	// Create a new User struct and populate it with form values
	newUser := User{
		Username: r.FormValue("username"),
		Email:    r.FormValue("email"),
		Password: r.FormValue("password"),
	}

	// Perform validation and save to DB

	// Respond with the created user
	w.Header().Set("Content-Type", "text/plain")
	fmt.Fprintf(w, "User created successfully:\nUsername: %s\nEmail: %s", newUser.Username, newUser.Email)
}

func handleNil(err error, errMsg string) {
	if err != nil {
		fmt.Println(errMsg)
		panic(err)
	}
}
