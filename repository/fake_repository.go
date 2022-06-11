package repository

import (
	"errors"
	"fmt"

	"github.com/sadovam/todos/models"
)

func MakeFakeData(totalLists, itemsPerList int) []*models.TodoList {
	todoLists := make([]*models.TodoList, totalLists, totalLists)
	for i := 0; i < totalLists; i++ {
		todos := make([]*models.TodoItem, itemsPerList, itemsPerList)
		for j := 0; j < itemsPerList; j++ {
			uid := i*itemsPerList + j
			todos[j] = &models.TodoItem{
				Uid:     uid,
				ListUid: i,
				Title:   fmt.Sprintf("<=Item %d List %d=>", uid, i),
				IsDone:  uid%3 < 2,
			}
		}
		todoLists[i] = &models.TodoList{
			Uid:   i,
			Title: fmt.Sprintf("<<<List %d>>>", i),
			Todos: todos,
		}
	}
	return todoLists
}

type TodosFake struct {
	data []*models.TodoList
}

func NewTodosFakeRepository(dataset []*models.TodoList) *TodosFake {
	return &TodosFake{data: dataset}
}

func (repo *TodosFake) GetItemByUid(uid int) (*models.TodoItem, error) {
	for _, list := range repo.data {
		for _, item := range list.Todos {
			if item.Uid == uid {
				return item, nil
			}
		}
	}
	return nil, errors.New(fmt.Sprintf("TodoItem with uid: %d doesn't exist", uid))
}

func (repo *TodosFake) GetListByUid(listUid int) (*models.TodoList, error) {
	for _, list := range repo.data {
		if listUid == list.Uid {
			return list, nil
		}
	}

	return nil, errors.New(fmt.Sprintf("List with uid: %d doesn't exist", listUid))
}

func (repo *TodosFake) findMaxItemUid() int {
	maxUid := 0

	for _, list := range repo.data {
		for _, item := range list.Todos {
			if item.Uid > maxUid {
				maxUid = item.Uid
			}
		}
	}

	return maxUid
}

func (repo *TodosFake) findMaxListUid() int {
	maxUid := 0
	for _, list := range repo.data {
		if list.Uid > maxUid {
			maxUid = list.Uid
		}
	}
	return maxUid
}

func (repo *TodosFake) CreateItem(listUid int, title string) (*models.TodoItem, error) {
	list, err := repo.GetListByUid(listUid)
	if err != nil {
		return nil, err
	}

	maxUid := repo.findMaxItemUid()
	item := &models.TodoItem{
		Uid:     maxUid + 1,
		ListUid: listUid,
		Title:   title,
		IsDone:  false,
	}
	list.Todos = append(list.Todos, item)

	return item, nil
}

func (repo *TodosFake) CreateList(title string) (*models.TodoList, error) {
	if title == "" {
		return nil, errors.New("Title can't be empty")
	}

	maxUid := repo.findMaxListUid()
	list := &models.TodoList{
		Uid:   maxUid + 1,
		Title: title,
		Todos: make([]*models.TodoItem, 0),
	}
	repo.data = append(repo.data, list)

	return list, nil
}

func (repo *TodosFake) UpdateItem(itemUid, listUid int, title string, isDone bool) (*models.TodoItem, error) {
	item, err := repo.GetItemByUid(itemUid)

	if err != nil {
		return nil, err
	}

	oldList, _ := repo.GetListByUid(item.ListUid)

	newList, err := repo.GetListByUid(listUid)
	if err != nil {
		return nil, err
	}

	item.ListUid = listUid
	item.Title = title
	item.IsDone = isDone

	newList.Todos = append(newList.Todos, item)
	for i, v := range oldList.Todos {
		if item.Uid == v.Uid {
			oldList.Todos = append(oldList.Todos[:i], oldList.Todos[i+1:]...)
		}
	}

	return item, nil
}

func (repo *TodosFake) testOfConsistency() bool {
	for _, list := range repo.data {
		for _, item := range list.Todos {
			if item.ListUid != list.Uid {
				return false
			}
		}
	}
	return true
}

func (repo *TodosFake) UpdateList(listUid int, title string) (*models.TodoList, error) {
	list, err := repo.GetListByUid(listUid)
	if err != nil {
		return nil, err
	}

	list.Title = title
	return list, nil
}

func (repo *TodosFake) DeleteItem(uid int) (*models.TodoItem, error) {
	for _, list := range repo.data {
		for i, item := range list.Todos {
			if item.Uid == uid {
				list.Todos = append(list.Todos[:i], list.Todos[i+1:]...)
				return item, nil
			}
		}
	}

	return nil, errors.New(fmt.Sprintf("TodoItem with uid: %d doesn't exist", uid))
}

func (repo *TodosFake) DeleteList(listUid int) (*models.TodoList, error) {
	for i, list := range repo.data {
		if list.Uid == listUid {
			if len(list.Todos) > 0 {
				return nil, errors.New(fmt.Sprintf("TodoList with uid: %d doesn't empty", listUid))
			}
			repo.data = append(repo.data[:i], repo.data[i+1:]...)
			return list, nil
		}
	}

	return nil, errors.New(fmt.Sprintf("TodoList with uid: %d doesn't exist", listUid))
}
