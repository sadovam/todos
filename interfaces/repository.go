package interfaces

import "github.com/sadovam/todos/models"

type TodosRepository interface {
	GetItemByUid(uid int) (*models.TodoItem, error)
	GetListByUid(uid int) (*models.TodoList, error)
	CreateItem(listUid int, title string) (*models.TodoItem, error)
	CreateList(title string) (*models.TodoList, error)
	UpdateItem(itemUid, listUid int, title string, isDone bool) (*models.TodoItem, error)
	UpdateList(listUid int, title string) (*models.TodoList, error)
	DeleteItem(uid int) (*models.TodoItem, error)
	DeleteList(listUid int) (*models.TodoList, error)
}
