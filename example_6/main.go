package main

import (
	"github.com/gorilla/mux"
	"labix.org/v2/mgo"
	"log"
	"net/http"
)

var (
	mongoSession *mgo.Session
	database     *mgo.Database
	repo         todoRepo
)

func main() {
	var err error

	// Setup the database
	if mongoSession, err = mgo.Dial("localhost"); err != nil {
		panic(err)
	}
  log.Println("Connected to mongodb")

	database = mongoSession.DB("mgo_examples_06")
	repo.Collection = database.C("todos")

	// Setup the web server and handlers
	r := mux.NewRouter()

	r.HandleFunc("/todos/{id}/complete", handleTodoComplete)
	r.HandleFunc("/todos/{id}/uncomplete", handleTodoUncomplete)
	r.HandleFunc("/todos/{id}", handleTodoDestroy).Methods("DELETE")
	r.HandleFunc("/todos/{id}", handleTodoUpdate).Methods("PUT")
	r.HandleFunc("/todos", handleTodoCreate).Methods("POST")
	r.HandleFunc("/todos", handleTodos).Methods("GET")
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Processing index.html: %v\n", r.RequestURI)
		http.ServeFile(w, r, "./index.html")
	})

	http.Handle("/", r)

	go func() {
    log.Printf("Starting webserver http://localhost:8080")
		panic(http.ListenAndServe(":8080", nil))
	}()
	<-make(chan bool)

}
