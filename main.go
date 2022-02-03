package main

import (
	"log"
	"net/http"
	"os"

	"github.com/Max-Gabriel-Susman/GoMicroservice/handlers"
)

func main() {

	// create log
	log := log.New(os.Stdout, "product-api", log.LstdFlags)

	// inject log into new handler
	helloHandler := handlers.NewHello(log)

	// create a new serve mux
	serveMux := http.NewServeMux()
	serveMux.Handle("/", helloHandler)

	// http server
	http.ListenAndServe(":9090", nil)
}
