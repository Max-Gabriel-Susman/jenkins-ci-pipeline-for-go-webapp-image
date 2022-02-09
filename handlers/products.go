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

	if request.Method == http.MethodPost {
		products.addProduct(responseWriter, request)
		return
	}
	// catch all
	responseWriter.WriteHeader(http.StatusMethodNotAllowed)
}

func (products *Products) getProducts(responseWriter http.ResponseWriter, request *http.Request) {
	productList := data.GetProducts()
	JSONData, err := json.Marshal(productList)
	if err != nil {
		http.Error(responseWriter, "Unable to marshal JSON.", http.StatusInternalServerError)
	}

	responseWriter.Write(JSONData)
}

func (p *Products) addProduct(rw http.ResponseWriter, r *http.Request) {
	p.logger.Println("Handle POST Products")

	prod := &data.Product{}

	err := prod.FromJSON(r.Body)
	if err != nil {
		http.Error(rw, "Unable to unmarshal json", http.StatusBadRequest)
	}

	p.logger.Printf("Prod: %#v", prod)
}
