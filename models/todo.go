package models

type TodoItem struct {
	Uid     int
	ListUid int
	Title   string
	IsDone  bool
}

type TodoList struct {
	Uid   int
	Title string
	Todos []TodoItem
}
