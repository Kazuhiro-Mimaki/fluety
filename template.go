package main

import (
	"html/template"
	"net/http"
)

type Template struct{}

func NewTemplate() *Template {
	return &Template{}
}

func (t *Template) Render() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		template, err := template.ParseFiles("index.html")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if err := template.Execute(w, "no data needed"); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}
