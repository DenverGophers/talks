package main

import (
	"fmt"
	"log"
	"net/http"
)

// START OMIT
func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello, web")
	})

	fmt.Println("Starting server")
	err := http.ListenAndServe("localhost:8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}

// END OMIT
