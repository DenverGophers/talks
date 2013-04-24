package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

// readJson will parses the JSON-encoded data in the http request and store the result in v
func readJson(r *http.Request, v interface{}) bool {
	defer r.Body.Close()

	var (
		body []byte
		err  error
	)

	body, err = ioutil.ReadAll(r.Body)

	if err != nil {
		log.Printf("ReadJson couldn't read request body %v", err)
		return false
	}

	if err = json.Unmarshal(body, v); err != nil {
		log.Printf("ReadJson couldn't parse request body %v", err)
		return false
	}

	return true
}
