package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func handleTodos(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		handleTodoPost(w, r)
	case "PUT":
		handleTodoPut(w, r)
	default:
		handleTodoGet(w, r)
	}
}

func handleTodoGet(w http.ResponseWriter, r *http.Request) {
	log.Printf("Processing TodoGet: %v\n", r.RequestURI)
	var (
		todos Todos
		err   error
	)
	if todos, err = repo.All(); err != nil {
		log.Printf("%v", err)
		http.Error(w, "500 Internal Server Error", 500)
		return
	}
	writeJson(w, todos)
}

func handleTodoPost(w http.ResponseWriter, r *http.Request) {}
func handleTodoPut(w http.ResponseWriter, r *http.Request)  {}

func handleTodoComplete(w http.ResponseWriter, r *http.Request) {
	var (
		todo Todo
		err  error
	)
	log.Printf("Processing TodoComplete: %v\n", r.RequestURI)
	data := struct {
		Success bool `json:"success"`
		Todo    Todo `json:"todo"`
	}{
		Success: false,
	}
	vars := mux.Vars(r)
	id := vars["id"]
	if todo, err = repo.Complete(id); err != nil {
		log.Printf("%v", err)
	} else {
		data.Success = true
		data.Todo = todo
	}

	writeJson(w, data)
}

func handleTodoUncomplete(w http.ResponseWriter, r *http.Request) {
	var (
		todo Todo
		err  error
	)
	log.Printf("Processing TodoUncomplete: %v\n", r.RequestURI)
	data := struct {
		Success bool `json:"success"`
		Todo    Todo `json:"todo"`
	}{
		Success: false,
	}
	vars := mux.Vars(r)
	id := vars["id"]
	if todo, err = repo.Uncomplete(id); err != nil {
		log.Printf("%v", err)
	} else {
		data.Success = true
		data.Todo = todo
	}

	writeJson(w, data)
}

func handleTodoDestroy(w http.ResponseWriter, r *http.Request) {
	var (
		err error
	)
	log.Printf("Processing TodoDestroy: %v\n", r.RequestURI)
	data := struct {
		Success bool `json:"success"`
	}{
		Success: false,
	}
	vars := mux.Vars(r)
	id := vars["id"]
	if err = repo.Destroy(id); err != nil {
		log.Printf("%v", err)
	} else {
		data.Success = true
	}

	writeJson(w, data)
}

func handleTodoCreate(w http.ResponseWriter, r *http.Request) {
	var (
		todo Todo
		err  error
	)
	log.Printf("Processing TodoCreate: %v\n", r.RequestURI)
	data := struct {
		Success bool `json:"success"`
		Todo    Todo `json:"todo"`
	}{
		Success: false,
	}
	if readJson(r, &todo) {
		if err = repo.Create(&todo); err != nil {
			log.Printf("%v", err)
		} else {
			data.Success = true
			data.Todo = todo
		}
	}

	writeJson(w, data)
}

func handleTodoUpdate(w http.ResponseWriter, r *http.Request) {
	var (
		todo Todo
		err  error
	)
	log.Printf("Processing TodoUpdate: %v\n", r.RequestURI)
	data := struct {
		Success bool `json:"success"`
		Todo    Todo `json:"todo"`
	}{
		Success: false,
	}
	if readJson(r, &todo) {
		if err = repo.Update(&todo); err != nil {
			log.Printf("%v", err)
		} else {
			data.Success = true
			data.Todo = todo
		}
	}

	writeJson(w, data)
}
