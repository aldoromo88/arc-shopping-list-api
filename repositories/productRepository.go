package repositories

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"

	models "github.com/aldoromo88/arc-shopping-list-api/models"
)

var products []models.Product

func init() {
	products = append(products, models.Product{ID: "1", Name: "Product One", Price: 1100})
	products = append(products, models.Product{ID: "2", Name: "Product Two", Price: 2200})
	products = append(products, models.Product{ID: "3", Name: "Product Three", Price: 3300})
}

// GetProducts : Get all products by search params
func GetProducts(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(products)
}

// GetProduct : Get product by ID
func GetProduct(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for _, item := range products {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&models.Product{})
}

// CreateProduct : Create a new product
func CreateProduct(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var product models.Product
	_ = json.NewDecoder(r.Body).Decode(&product)
	product.ID = params["id"]
	products = append(products, product)
	json.NewEncoder(w).Encode(products)

}

// DeleteProduct : Detele product by ID
func DeleteProduct(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for index, item := range products {
		if item.ID == params["id"] {
			products = append(products[:index], products[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(products)
}
