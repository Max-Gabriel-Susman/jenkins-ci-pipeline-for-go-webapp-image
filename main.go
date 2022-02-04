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

	// create log
	log := log.New(os.Stdout, "product-api", log.LstdFlags)

	// inject log into new handler
	helloHandler := handlers.NewHello(log)
	goodbyeHandler := handlers.NewGoodbye(log)

	// create a new serve mux and register hello and goodbye handlers
	serveMux := http.NewServeMux()
	serveMux.Handle("/", helloHandler)
	serveMux.Handle("/goodbye", goodbyeHandler)

	// create server
	server := &http.Server{
		Addr:         ":9090",
		Handler:      serveMux,
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}

	go func() {
		// listen and serve on TCP
		err := server.ListenAndServe()
		if err != nil {
			log.Fatal(err)
		}
	}()

	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, os.Kill)

	sig := <-sigChan
	log.Println("Received terminate, graceful shutdown", sig)
	// create deadline context
	timeoutContext, _ := context.WithTimeout(context.Background(), 30*time.Second)

	// graceful shutdown
	server.Shutdown(timeoutContext)
}
