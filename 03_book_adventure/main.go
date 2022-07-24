package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/mathiashandle/03adventure/story"
)

func main() {
	port := flag.Int("port", 3001, "port to start the server on")
	fileName := flag.String("file", "data.json", "the JSON file with CYOA story")
	flag.Parse()

	f, err := os.Open(*fileName)
	if err != nil {
		log.Fatalf("couldn't open a file %s", err)
	}
	s := story.ParseJSON(f)

	h := story.NewHandler(s)
	fmt.Printf("Starting the server at: %d\n", *port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *port), h))
}
