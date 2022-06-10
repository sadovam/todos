package repository

import (
	"errors"
	"fmt"

	"github.com/sadovam/todos/models"
)

func MakeFakeData(totalLists, itemsPerList int) []models.TodoList {
	todoLists := make([]models.TodoList, totalLists, totalLists)
	for i := 0; i < totalLists; i++ {
		todos := make([]models.TodoItem, itemsPerList, itemsPerList)
		for j := 0; j < itemsPerList; j++ {
			uid := i*itemsPerList + j
			todos[j] = models.TodoItem{
				Uid:     uid,
				ListUid: i,
				Title:   fmt.Sprintf("<=Item %d List %d=>", uid, i),
				IsDone:  uid%3 < 2,
			}
		}
		todoLists[i] = models.TodoList{
			Uid:   i,
			Title: fmt.Sprintf("<<<List %d>>>", i),
			Todos: todos,
		}
	}
	return todoLists
}

type TodosFake struct {
	data []models.TodoList
}

func NewTodosFakeRepository(dataset []models.TodoList) TodosFake {
	return TodosFake{data: dataset}
}

func (repo TodosFake) GetItemByUid(uid int) (*models.TodoItem, error) {
	for _, list := range repo.data {
		for _, item := range list.Todos {
			if item.Uid == uid {
				return &item, nil
			}
		}
	}
	return nil, errors.New(fmt.Sprintf("TodoItem with uid: %d doesn't exist", uid))
}

func (repo TodosFake) GetListByUid(listUid int) (*models.TodoList, error) {
	for _, list := range repo.data {
		if listUid == list.Uid {
			return &list, nil
		}
	}

	return nil, errors.New(fmt.Sprintf("List with uid: %d doesn't exist", listUid))
}
