package interfaces

import "github.com/sadovam/todos/models"

type TodosRepository interface {
	GetItemByUid(uid int) (models.TodoItem, error)
	GetItemsByListUid(listUid int) ([]models.TodoItem, error)
}
