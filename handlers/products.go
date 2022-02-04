package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/Max-Gabriel-Susman/GoMicroservice/data"
)

type Products struct {
	logger *log.Logger
}

func NewProducts(logger *log.Logger) *Products {
	return &Products{logger}
}

func (products *Products) ServeHTTP(responseWriter http.ResponseWriter, request *http.Request) {
	if request.Method == http.MethodGet {
		products.getProducts(responseWriter, request)
		return
	}
}

func (products *Products) getProducts(responseWriter http.ResponseWriter, request *http.Request) {
	productList := data.GetProducts()
	JSONData, err := json.Marshal(productList)
	if err != nil {
		http.Error(responseWriter, "Unable to marshal JSON.", http.StatusInternalServerError)
	}

	responseWriter.Write(JSONData)
}
