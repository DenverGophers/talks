package main

import (
	"html/template"
	"log"
	"net/http"
)

// one OMIT
const bind = ":8000"

func main() {
	http.Handle("/", tplHandler(Common))
	http.Handle("/a", tplHandler(VariantA))
	http.Handle("/b", tplHandler(VariantB))
	log.Println("Serving on", bind)
	if err := http.ListenAndServe(bind, nil); err != nil {
		log.Fatalln(err)
	}
}

func tplHandler(t *template.Template) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("Handling", r.URL, "with template", t.Name())
		if err := t.Execute(w, nil); err != nil {
			log.Println(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})
}

// two OMIT

var Common, VariantA, VariantB *template.Template

func init() {
	tm := template.Must
	Common = tm(template.New("main").Parse(`
		<!DOCTYPE html>
		<html>
			<head>
			{{template "head"}}
			</head>
			<body>
			<p>Something common. See /a and /b</p>
			{{template "body"}}
			</body>
		</html>
		{{/* these provide empty defaults */}}
		{{define "head"}}{{end}}
		{{define "body"}}{{end}}
	`))
	// three OMIT
	VariantA = tm(Common.Clone())
	tm(VariantA.New("body").Parse(`
		{{define "head"}}
		<title>Variant A</title>
		{{end}}
		{{/* what follows is implicitly the "body" template */}}
		Body for Variant A
	`))
	VariantB = tm(Common.Clone())
	tm(VariantB.New("body").Parse(`
		{{define "head"}}
		<title>Variant B</title>
		{{end}}
		{{/* what follows is implicitly the (other) "body" template */}}
		Variant B Body
	`))
}
