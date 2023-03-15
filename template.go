package main

import (
	"net/http"
	"text/template"
)

type Todo struct {
	Title string
	Done bool
}

type TodoPageData struct {
	PageTitle string
	Todos []Todo
}

func ServeTemplate() {
	// tmpl, err := template.ParseFiles("static/layout.html")
	// or
	tmpl := template.Must(template.ParseFiles("static/layout.html"))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		data := TodoPageData{
				PageTitle: "My TODO list",
				Todos: []Todo{
						{Title: "Task 1", Done: false},
						{Title: "Task 2", Done: true},
						{Title: "Task 3", Done: true},
				},
		}

		tmpl.Execute(w, data)
	})

	http.ListenAndServe(":80", nil)
}