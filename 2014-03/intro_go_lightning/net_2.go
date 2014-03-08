package main

import (
	"fmt"
	"log"
	"net/http"
)

// START OMIT
func main() {
	http.HandleFunc("/", handler)
	log.Println("Starting server on: localhost:8080")
	log.Println("localhost:8080")
	err := http.ListenAndServe("localhost:8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello, web")
}

// END OMIT
