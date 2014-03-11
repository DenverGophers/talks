package main

import (
	"io"
	"log"
	"net/http"
	"os"
)

// START OMIT
func main() {
	r, err := http.Get("http://www.golang.org/")
	if err != nil {
		log.Fatal(err)
	}
	if r.StatusCode != http.StatusOK {
		log.Fatal(r.Status)
	}
	io.Copy(os.Stdout, r.Body)
}

// END OMIT
