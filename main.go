package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"github.com/sadovam/todos/components/page"
	"github.com/sadovam/todos/components/todo"
	"github.com/sadovam/todos/repository"
	"github.com/sadovam/todos/services"
)

type htmlResponse struct {
	Error string `json:"error"`
	Node  string `json:"node"`
}

var data = repository.MakeFakeData(3, 5)
var repo = repository.NewTodosFakeRepository(data)
var serv = services.NewTodoItemService(repo)

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "<h1>Hello</h1>")
}

func deleteTodo(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	uid, _ := strconv.Atoi(vars["uid"])
	serv.DeleteItem(uid)
}

func createTodo(w http.ResponseWriter, r *http.Request) {
	title := r.FormValue("title")
	error, item := "", ""
	data, err := serv.CreateItem(1, title)
	if err != nil {
		error = err.Error()
	} else {
		item = todo.TodoItem(data)
	}
	resp := htmlResponse{Error: error, Node: item}
	json.NewEncoder(w).Encode(resp)
}

func getTodos(w http.ResponseWriter, r *http.Request) {
	data, _ := serv.GetTodoListByUid(1)
	fmt.Fprintf(w, page.Page(todo.TodoList(data)))
}

func main() {
	fmt.Println("Hello todos!")

	r := mux.NewRouter()

	fs := http.FileServer(http.Dir("static/"))
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static", fs))
	r.HandleFunc("/todo", getTodos).Methods("GET")
	r.HandleFunc("/todo", createTodo).Methods("POST")
	r.HandleFunc("/todo/{uid}", deleteTodo).Methods("DELETE")

	r.HandleFunc("/", index)

	port := 3390
	fmt.Printf("Starting on port %d...\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), r))
}
