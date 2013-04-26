package main

import (
	"github.com/gorilla/mux"
	"labix.org/v2/mgo"
	"log"
	"net/http"
	"time"
)

var (
	mongoSession *mgo.Session
	database     *mgo.Database
	repo         todoRepo

	router = mux.NewRouter()
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
	// START OMIT
	route("/todos/{id}/complete", handleTodoComplete)
	route("/todos/{id}/complete", handleTodoComplete)
	route("/todos/{id}/uncomplete", handleTodoUncomplete)
	route("/todos/{id}", handleTodoDestroy).Methods("DELETE")
	route("/todos/{id}", handleTodoUpdate).Methods("PUT")
	route("/todos", handleTodoCreate).Methods("POST")
	route("/todos", handleTodos).Methods("GET")
	route("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./index.html")
	})
	// END OMIT

	http.Handle("/", router)

	log.Printf("Starting webserver http://localhost:8080")
	panic(http.ListenAndServe(":8080", nil))
}

func route(pattern string, handler func(http.ResponseWriter, *http.Request)) *mux.Route {
	handler = logRequest(handler)
	return router.HandleFunc(pattern, handler)
}

func logRequest(handler func(http.ResponseWriter, *http.Request)) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var s = time.Now()
		handler(w, r)
		log.Printf("%s %s %6.3fms", r.Method, r.RequestURI, (time.Since(s).Seconds()*1000))
	}
}
