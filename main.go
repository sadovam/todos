package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"text/template"

	"github.com/gorilla/mux"

	"github.com/sadovam/todos/models"
	"github.com/sadovam/todos/repository"
	"github.com/sadovam/todos/services"
)

var data = repository.MakeFakeData(3, 5)
var repo = repository.NewTodosFakeRepository(data)
var serv = services.NewTodoItemService(repo)

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "<h1>Hello</h1>")
}

func createTodo(w http.ResponseWriter, r *http.Request) {
	title := r.FormValue("title")
	tmpl := template.Must(template.ParseFiles("templates/components/todo.html"))
	data := models.TodoItem{IsDone: false, Title: title}
	tmpl.Execute(w, data)
}

func todo(w http.ResponseWriter, r *http.Request) {
	root := "templates/todo/"
	fileSystem := os.DirFS(root)
	tmpl := template.Must(template.ParseFS(fileSystem, "*.html"))
	data, _ := serv.GetTodoListByUid(1)
	tmpl.ExecuteTemplate(w, "page", data)
}

func main() {
	fmt.Println("Hello todos!")

	r := mux.NewRouter()

	fs := http.FileServer(http.Dir("static/"))
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static", fs))
	r.HandleFunc("/todo", todo).Methods("GET")
	r.HandleFunc("/todo", createTodo).Methods("POST")

	r.HandleFunc("/", index)

	port := 3390
	fmt.Printf("Starting on port %d...\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), r))
}
