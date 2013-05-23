package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	var count int

	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		count++
		io.WriteString(w, fmt.Sprintf("demo 2; hello %d\n", count))
	})

	if err := http.ListenAndServe(":4001", nil); err != nil {
		panic(err)
	}
}
