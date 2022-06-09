package services

import (
	"github.com/sadovam/todos/interfaces"
	"github.com/sadovam/todos/models"
)

type TodoItemService struct {
	repo interfaces.TodosRepository
}

func NewTodoItemService(repository interfaces.TodosRepository) TodoItemService {
	return TodoItemService{repo: repository}
}

func (serv TodoItemService) GetTodoItemByUId(uid int) (models.TodoItem, error) {
	return serv.repo.GetItemByUid(uid)
}

func (serv TodoItemService) GetTodoItemsByListUid(listUid int) ([]models.TodoItem, error) {
	return serv.repo.GetItemsByListUid(listUid)
}
