package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/Max-Gabriel-Susman/GoMicroservice/handlers"
)

func main() {

	// create logger
	logger := log.New(os.Stdout, "product-api", log.LstdFlags)

	// inject log into new handler
	productHandler := handlers.NewProducts(logger)
	// create a new serve mux and register hello and goodbye handlers
	serveMux := http.NewServeMux()
	serveMux.Handle("/", productHandler)

	// create server
	server := &http.Server{
		Addr:         ":9090",
		Handler:      serveMux,
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}

	// run as a background process
	go func() {
		// listen and serve on TCP
		err := server.ListenAndServe()
		if err != nil {
			logger.Fatal(err)
		}
	}()

	// create an os signal
	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, os.Kill)

	// sig is assigned (gotta grok arrow operator)
	sig := <-sigChan
	logger.Println("Received terminate, graceful shutdown", sig)
	// create deadline context
	timeoutContext, _ := context.WithTimeout(context.Background(), 30*time.Second)
	// what is a context leak?

	// graceful shutdown
	server.Shutdown(timeoutContext)
}
