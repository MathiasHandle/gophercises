package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/mathiashandle/03adventure/handlers"
	"github.com/mathiashandle/03adventure/helpers"
)

func main() {
	fileName := flag.String("file", "data.json", "the JSON file with CYOA story")
	flag.Parse()

	f, err := os.Open(*fileName)
	if err != nil {
		log.Fatalf("couldn't open a file %s", err)
	}
	helpers.ParseJSON(f)

	mux := http.NewServeMux()
	server := http.Server{
		Addr:    ":3001",
		Handler: mux,
	}

	mux.HandleFunc("/", handlers.Homepage)

	fmt.Printf("Starting up server on port %s \n", server.Addr)
	err = server.ListenAndServe()
	if err != nil {
		fmt.Println("Error starting server: ", err)
	}
}
