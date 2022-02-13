package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"fmt"
	"github.com/Max-Gabriel-Susman/GoMicroservice/data"
	"github.com/gorilla/mux"
)

type Products struct {
	logger *log.Logger
}

func NewProducts(logger *log.Logger) *Products {
	return &Products{logger}
}

func (products *Products) GetProducts(responseWriter http.ResponseWriter, request *http.Request) {
	productList := data.GetProducts()
	JSONData, err := json.Marshal(productList)
	if err != nil {
		http.Error(responseWriter, "Unable to marshal JSON.", http.StatusInternalServerError)
	}

	responseWriter.Write(JSONData)
}

func (p *Products) AddProduct(rw http.ResponseWriter, r *http.Request) {
	p.logger.Println("Handle POST Products")

	prod := &data.Product{}

	err := prod.FromJSON(r.Body)
	if err != nil {
		http.Error(rw, "Unable to unmarshal json", http.StatusBadRequest)
	}

	data.AddProduct(prod)
}

func (p *Products) UpdateProducts(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(rw, "Unable to convert id", http.StatusBadRequest)
		return 
	}

	p.logger.Println("Handle PUT Product", id)

	err = data.UpdateProduct(id, prod)
	if err == data.ErrProductNotFound {
		fmt.Println("update product")
		http.Error(rw, "Product not found", http.StatusNotFound)
		return
	}

	if err != nil {
		http.Error(rw, "Product not found", http.StatusInternalServerError)
		return
	}
	p.logger.Printf("Prod: %#v", prod)
	data.UpdateProduct(id, prod)
}

type KeyProduct struct

func (p Products) MiddlewareProductValidation(next http.Handler) http.Handler {
	return http.HandleFunc(rw http.ResponseWriter, r *http.Request) {
		prod := &data.Product{}

		err := prod.FromJSON(r.Body)
		if err != nil {
			http.Error(rw, "Unable to marshal JSON", http.StatusBadRequest)
			return 
		}

		ctx := r.Context().WithValue(KeyProduct{}, prod)
		req := r.WithContext(ctx)

		next.ServeHTTP(rw, req)
	}
}