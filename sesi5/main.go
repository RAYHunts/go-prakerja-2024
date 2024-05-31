package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"
)

type Product struct {
	ID    int
	Name  string
	Price int
}

var products []Product

func main() {
	// Read JSON file
	file, err := os.ReadFile("products.json")
	if err != nil {
		fmt.Println(err)
		return
	}

	err = json.Unmarshal(file, &products)
	if err != nil {
		fmt.Println(err)
		return
	}

	http.HandleFunc("/products", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			// Return all products
			productsJSON, err := json.Marshal(products)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			w.Header().Set("Content-Type", "application/json")
			w.Write(productsJSON)
		case "POST":
			// Create a new product
			var newProduct Product
			err := json.NewDecoder(r.Body).Decode(&newProduct)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			defer r.Body.Close()

			// Generate a new ID for the product
			if len(products) > 0 {
				newProduct.ID = products[len(products)-1].ID + 1
			} else {
				newProduct.ID = 1
			}

			products = append(products, newProduct)

			// Rewrite JSON file
			err = rewriteJSONFile()
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			w.WriteHeader(http.StatusCreated)
			w.Write([]byte("Product created"))
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	http.HandleFunc("/products/", func(w http.ResponseWriter, r *http.Request) {
		// Extract the product ID from the URL
		id := strings.TrimPrefix(r.URL.Path, "/products/")
		idInt, err := strconv.Atoi(id)
		if err != nil {
			http.Error(w, "Invalid product ID", http.StatusBadRequest)
			return
		}

		switch r.Method {
		case "GET":
			// Return a product by ID
			for _, product := range products {
				if product.ID == idInt {
					productJSON, err := json.Marshal(product)
					if err != nil {
						http.Error(w, err.Error(), http.StatusInternalServerError)
						return
					}
					w.Header().Set("Content-Type", "application/json")
					w.Write(productJSON)
					return
				}
			}
			http.Error(w, "Product not found", http.StatusNotFound)
		case "PUT":
			// Update a product by ID
			var updatedProduct Product
			err := json.NewDecoder(r.Body).Decode(&updatedProduct)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			defer r.Body.Close()

			for i, product := range products {
				if product.ID == idInt {
					updatedProduct.ID = product.ID
					products[i] = updatedProduct

					// Rewrite JSON file
					err = rewriteJSONFile()
					if err != nil {
						http.Error(w, err.Error(), http.StatusInternalServerError)
						return
					}

					w.Write([]byte("Product updated"))
					return
				}
			}
			http.Error(w, "Product not found", http.StatusNotFound)
		case "DELETE":
			// Delete a product by ID
			for i, product := range products {
				if product.ID == idInt {
					products = append(products[:i], products[i+1:]...)

					// Rewrite JSON file
					err = rewriteJSONFile()
					if err != nil {
						http.Error(w, err.Error(), http.StatusInternalServerError)
						return
					}

					w.Write([]byte("Product deleted"))
					return
				}
			}
			http.Error(w, "Product not found", http.StatusNotFound)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}

func rewriteJSONFile() error {
	productsJSON, err := json.Marshal(products)
	if err != nil {
		return err
	}

	err = os.WriteFile("products.json", productsJSON, 0644)
	if err != nil {
		return err
	}

	return nil
}
