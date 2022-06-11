package services

import (
	"github.com/sadovam/todos/interfaces"
	"github.com/sadovam/todos/models"
)

type TodoItemService struct {
	repo interfaces.TodosRepository
}

func NewTodoItemService(repository interfaces.TodosRepository) *TodoItemService {
	return &TodoItemService{repo: repository}
}

func (serv *TodoItemService) GetTodoItemByUId(uid int) (*models.TodoItem, error) {
	return serv.repo.GetItemByUid(uid)
}

func (serv *TodoItemService) GetTodoListByUid(uid int) (*models.TodoList, error) {
	return serv.repo.GetListByUid(uid)
}

func (serv *TodoItemService) CreateItem(listUid int, title string) (*models.TodoItem, error) {
	return serv.repo.CreateItem(listUid, title)
}
func (serv *TodoItemService) CreateList(title string) (*models.TodoList, error) {
	return serv.repo.CreateList(title)
}
func (serv *TodoItemService) UpdateItem(itemUid, listUid int, title string, isDone bool) (*models.TodoItem, error) {
	return serv.repo.UpdateItem(itemUid, listUid, title, isDone)
}
func (serv *TodoItemService) UpdateList(listUid int, title string) (*models.TodoList, error) {
	return serv.repo.UpdateList(listUid, title)
}
func (serv *TodoItemService) DeleteItem(uid int) (*models.TodoItem, error) {
	return serv.repo.DeleteItem(uid)
}
func (serv *TodoItemService) DeleteList(listUid int) (*models.TodoList, error) {
	return serv.repo.DeleteList(listUid)
}
