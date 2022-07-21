package main

import (
	"fmt"
	"net/http"

	"github.com/mathiashandle/03adventure/handlers"
	"github.com/mathiashandle/03adventure/helpers"
)

func main() {
	mux := http.NewServeMux()

	server := http.Server{
		Addr:    ":3001",
		Handler: mux,
	}

	mux.HandleFunc("/", handlers.Homepage)

	helpers.ParseJSON("data.json")

	fmt.Printf("Starting up server on port %s \n", server.Addr)
	err := server.ListenAndServe()
	if err != nil {
		fmt.Println("Error starting server: ", err)
	}
}
