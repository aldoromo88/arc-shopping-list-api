package main

import (
	"log"
	"net/http"

	repos "github.com/aldoromo88/arc-shopping-list-api/repositories"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/api/product", repos.GetProducts).Methods("GET")
	router.HandleFunc("/api/product/{id}", repos.GetProduct).Methods("GET")
	router.HandleFunc("/api/product/{id}", repos.CreateProduct).Methods("POST")
	router.HandleFunc("/api/product/{id}", repos.DeleteProduct).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", router))
}
