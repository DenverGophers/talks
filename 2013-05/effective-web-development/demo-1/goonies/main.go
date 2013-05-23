package main

import (
	"io"
	"net/http"
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		io.WriteString(w, "hello, world!\n")
	})

	if err := http.ListenAndServe(":4000", nil); err != nil {
		panic(err)
	}
}
