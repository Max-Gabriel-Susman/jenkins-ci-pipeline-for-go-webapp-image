package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"regexp"
	"strconv"

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

	if request.Method == http.MethodPut {
		products.logger.Println("PUT", request.URL.Path)
		reg := regexp.MustCompile(`/([0-9]+)`)
		g := reg.FindAllStringSubmatch(request.URL.Path, -1)

		if len(g) != 1 {
			products.logger.Println("Invalid URI more than one id")
			http.Error(responseWriter, "Invalid URI", http.StatusBadRequest)
			return
		}

		if len(g[0]) != 2 {
			products.logger.Println("Invalid URI more than one capture group")
			http.Error(responseWriter, "Invalid URI", http.StatusBadRequest)
			return
		}

		idString := g[0][1]
		id, err := strconv.Atoi(idString)
		if err != nil {
			products.logger.Println("Invalid URI unable to convert to number", idString)
			http.Error(responseWriter, "Invalid URI", http.StatusBadRequest)
			return
		}

		products.updateProducts(id, responseWriter, request)
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

	data.AddProduct(prod)
}

func (p *Products) updateProducts(id int, rw http.ResponseWriter, r *http.Request) {
	p.logger.Println("Handle PUT Products")

	prod := &data.Product{}

	err := prod.FromJSON(r.Body)
	if err != nil {
		http.Error(rw, "Unable to unmarshal json", http.StatusBadRequest)
	}

	err = data.UpdateProduct(id, prod)
	if err == data.ErrProductNotFound {
		http.Error(rw, "Product not found", http.StatusNotFound)
		return
	}

	if err != nil {
		http.Error(rw, "Product not found", http.StatusNotFound)
		return
	}
	p.logger.Printf("Prod: %#v", prod)
}
