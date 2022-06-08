package models

type TodoItem struct {
	Title  string
	IsDone bool
}

type TodoList struct {
	Title string
	Todos []TodoItem
}
