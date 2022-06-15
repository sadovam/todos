package repository

import (
	"fmt"
	"testing"

	"github.com/sadovam/todos/models"
)

func TestMakeFakeData(t *testing.T) {
	data := MakeFakeData(3, 5)
	if len(data) != 3 {
		t.Fatalf("Generating fake data error; may create %d lists, but creating %d",
			3, len(data))
	}
	if len(data[0].Todos) != 5 {
		t.Fatalf("Generating fake data error; may create %d todos per list, but creating %d",
			5, len(data))
	}
}

func TestGetItemByUid(t *testing.T) {
	data := MakeFakeData(5, 7)
	repo := NewTodosFakeRepository(data)
	res, err := repo.GetItemByUid(23)

	if err != nil {
		t.Fatalf("GetItemByUid error: %v", err.Error())
	}

	if res.Uid != 23 {
		t.Fatalf("GetItemByUid error: want uid = %d, got uid = %d", 23, res.Uid)
	}

}

func TestGetItemByUidError(t *testing.T) {
	data := MakeFakeData(2, 4)
	repo := NewTodosFakeRepository(data)
	res, err := repo.GetItemByUid(15)

	if err == nil {
		t.Fatalf("GetItemByUid error: want error, got result uid = %d", res.Uid)
	}
	want := fmt.Sprintf("TodoItem with uid: %d doesn't exist", 15)
	if err.Error() != want {
		t.Fatalf("GetItemByUid error: want %s, got %s", want, err.Error())
	}
}

func TestGetListByUid(t *testing.T) {
	data := MakeFakeData(5, 2)
	repo := NewTodosFakeRepository(data)
	res, err := repo.GetListByUid(3)

	if err != nil {
		t.Fatalf("GetListByUid error: %v", err.Error())
	}

	if res.Uid != 3 {
		t.Fatalf("GetlistByUid error: want uid = %d, got uid = %d", 3, res.Uid)
	}
}

func TestGetListByUidError(t *testing.T) {
	data := MakeFakeData(2, 4)
	repo := NewTodosFakeRepository(data)
	res, err := repo.GetListByUid(5)

	if err == nil {
		t.Fatalf("GetListByUid error: want error, got result uid = %d", res.Uid)
	}
	want := fmt.Sprintf("List with uid: %d doesn't exist", 5)
	if err.Error() != want {
		t.Fatalf("GetListByUid error: want %s, got %s", want, err.Error())
	}
}

func TestFindMaxItemUid(t *testing.T) {
	data := MakeFakeData(5, 3)
	repo := NewTodosFakeRepository(data)
	res := repo.findMaxItemUid()

	if res != 14 {
		t.Fatalf("findMaxItemUid error: want %d, got %d", 14, res)
	}
}

func TestFindMaxListUid(t *testing.T) {
	data := MakeFakeData(7, 2)
	repo := NewTodosFakeRepository(data)
	res := repo.findMaxListUid()

	if res != 6 {
		t.Fatalf("findMaxItemUid error: want %d, got %d", 6, res)
	}
}

func TestCreateItem(t *testing.T) {
	data := MakeFakeData(3, 4)
	repo := NewTodosFakeRepository(data)
	got, err := repo.CreateItem(1, "New todo item")

	if err != nil {
		t.Fatalf("CreateItem error: %v", err.Error())
	}

	want := &models.TodoItem{Uid: 12, ListUid: 1, Title: "New todo item", IsDone: false}

	if !want.IsEqual(got) {
		t.Fatalf("CreateItem error: want %v, got %v", want, got)
	}

	inList, err := repo.GetItemByUid(12)

	if err != nil {
		t.Fatalf("CreateItem error while trying to get new item from list: %v", err.Error())
	}

	if !inList.IsSame(got) {
		t.Fatalf("CreateItem error: want %v, got %v", inList, got)
	}
}

func TestCreateItemError(t *testing.T) {
	data := MakeFakeData(3, 4)
	repo := NewTodosFakeRepository(data)
	got, err := repo.CreateItem(7, "New todo item")

	if err == nil {
		t.Fatalf("CreateItem error: want error, got result %v", got)
	}

	want := fmt.Sprintf("List with uid: %d doesn't exist", 7)
	if err.Error() != want {
		t.Fatalf("CreateItem error: want %s, got %s", want, err.Error())
	}

	got, err = repo.CreateItem(2, "")

	if err == nil {
		t.Fatalf("CreateItem error: want error, got result %v", got)
	}

	want = "Title can't be empty"
	if err.Error() != want {
		t.Fatalf("CreateList error: want %s, got %s", want, err.Error())
	}
}

func TestCreateList(t *testing.T) {
	data := MakeFakeData(3, 2)
	repo := NewTodosFakeRepository(data)
	got, err := repo.CreateList("New todo list")

	if err != nil {
		t.Fatalf("CreateList error: %v", err.Error())
	}

	want := &models.TodoList{Uid: 3, Title: "New todo list", Todos: make([]*models.TodoItem, 0)}

	if !want.IsEqual(got) {
		t.Fatalf("CreateList error: want %v, got %v", want, got)
	}

	inData, err := repo.GetListByUid(3)

	if err != nil {
		t.Fatalf("CreateList error while trying to get new item from list: %v", err.Error())
	}

	if !inData.IsSame(got) {
		t.Fatalf("CreateList error: want %v, got %v", inData, got)
	}
}

func TestCreateListError(t *testing.T) {
	data := MakeFakeData(3, 2)
	repo := NewTodosFakeRepository(data)
	got, err := repo.CreateList("")

	if err == nil {
		t.Fatalf("CreateList error: want error, got result %v", got)
	}

	want := "Title can't be empty"
	if err.Error() != want {
		t.Fatalf("CreateList error: want %s, got %s", want, err.Error())
	}
}

func TestUpdateItem(t *testing.T) {
	data := MakeFakeData(3, 4)
	repo := NewTodosFakeRepository(data)
	got, err := repo.UpdateItem(7, 0, "New todo item", false)

	if err != nil {
		t.Fatalf("UpdateItem error: %v", err.Error())
	}

	want := &models.TodoItem{Uid: 7, ListUid: 0, Title: "New todo item", IsDone: false}

	if !want.IsEqual(got) {
		t.Fatalf("UpdateItem error: want %v, got %v", want, got)
	}

	inList, err := repo.GetItemByUid(7)

	if err != nil {
		t.Fatalf("UpdateItem error while trying to get new item from list: %v", err.Error())
	}

	if !inList.IsSame(got) {
		t.Fatalf("UpdateItem error: want %v, got %v", inList, got)
	}

	if !repo.testOfConsistency() {
		t.Fatal("UpdateItem error: consistensy destroyed!")
	}
}

func TestUpdateItemError(t *testing.T) {
	data := MakeFakeData(3, 4)
	repo := NewTodosFakeRepository(data)
	got, err := repo.UpdateItem(3, 7, "New todo item", false)

	if err == nil {
		t.Fatalf("UpdateItem error: want error, got result %v", got)
	}

	want := fmt.Sprintf("List with uid: %d doesn't exist", 7)
	if err.Error() != want {
		t.Fatalf("UpdateItem error: want %s, got %s", want, err.Error())
	}

	got, err = repo.UpdateItem(17, 0, "New todo item", false)

	if err == nil {
		t.Fatalf("UpdateItem error: want error, got result %v", got)
	}

	want = fmt.Sprintf("TodoItem with uid: %d doesn't exist", 17)
	if err.Error() != want {
		t.Fatalf("UpdateItem error: want %s, got %s", want, err.Error())
	}
}

func TestTestOfConsistency(t *testing.T) {
	data := MakeFakeData(5, 3)
	repo := NewTodosFakeRepository(data)

	res := repo.testOfConsistency()
	if !res {
		t.Fatalf("TestOfConsistency error: want %t, got %t", true, false)
	}

	data[1].Todos[2].ListUid = 3

	res = repo.testOfConsistency()
	if res {
		t.Fatalf("TestOfConsistency error: want %t, got %t", false, true)
	}
}

func TestUpdateList(t *testing.T) {
	data := MakeFakeData(3, 2)
	repo := NewTodosFakeRepository(data)
	res, err := repo.UpdateList(1, "Updated List")

	if err != nil {
		t.Fatalf("TestUpdateList error: %v", err.Error())
	}

	if res.Title != "Updated List" || res.Uid != 1 {
		t.Fatalf("TestUpdateList error: want title = %s, uid = %d, got title = %s, uid = %d",
			"Updated List", 1, res.Title, res.Uid)
	}
	want, _ := repo.GetListByUid(1)
	if !want.IsSame(res) {
		t.Fatalf("TestUpdateList error while read updated list from data: want %v, got %v", want, res)
	}
}

func TestUpdateListError(t *testing.T) {
	data := MakeFakeData(3, 4)
	repo := NewTodosFakeRepository(data)
	got, err := repo.UpdateList(5, "Updated List")

	if err == nil {
		t.Fatalf("UpdateList error: want error, got result %v", got)
	}

	want := fmt.Sprintf("List with uid: %d doesn't exist", 5)
	if err.Error() != want {
		t.Fatalf("UpdateList error: want %s, got %s", want, err.Error())
	}
}

func TestDeleteItem(t *testing.T) {
	data := MakeFakeData(3, 4)
	repo := NewTodosFakeRepository(data)
	got, err := repo.DeleteItem(7)

	if err != nil {
		t.Fatalf("DeleteItem error: %v", err.Error())
	}

	want := &models.TodoItem{Uid: 7, ListUid: 1, Title: "Item 7 List 1", IsDone: true}

	if !want.IsEqual(got) {
		t.Fatalf("DeleteItem error: want %v, got %v", want, got)
	}

	inList, err := repo.GetItemByUid(7)

	if err == nil {
		t.Fatalf("DeleteItem error: while trying to get it from list want error, got: %v", inList)
	}
}

func TestDeleteItemError(t *testing.T) {
	data := MakeFakeData(2, 4)
	repo := NewTodosFakeRepository(data)
	res, err := repo.DeleteItem(15)

	if err == nil {
		t.Fatalf("DeleteItem error: want error, got result uid = %d", res.Uid)
	}
	want := fmt.Sprintf("TodoItem with uid: %d doesn't exist", 15)
	if err.Error() != want {
		t.Fatalf("DeleteItem error: want %s, got %s", want, err.Error())
	}
}

func TestDeleteList(t *testing.T) {
	data := MakeFakeData(3, 2)
	repo := NewTodosFakeRepository(data)
	repo.DeleteItem(4)
	repo.DeleteItem(5)
	got, err := repo.DeleteList(2)

	if err != nil {
		t.Fatalf("DeleteList error: %v", err.Error())
	}

	if got.Uid != 2 {
		t.Fatalf("DeleteList error: want list with uid %d, got %d", 2, got.Uid)
	}

	inData, err := repo.GetListByUid(2)

	if err == nil {
		t.Fatalf("DeleteList error: while trying to get it from data want error, got: %v", inData)
	}
}

func TestDeleteListError(t *testing.T) {
	data := MakeFakeData(2, 4)
	repo := NewTodosFakeRepository(data)
	res, err := repo.DeleteList(15)

	if err == nil {
		t.Fatalf("DeleteList error: want error, got result uid = %d", res.Uid)
	}
	want := fmt.Sprintf("TodoList with uid: %d doesn't exist", 15)
	if err.Error() != want {
		t.Fatalf("DeleteList error: want %s, got %s", want, err.Error())
	}

	res, err = repo.DeleteList(0)
	if err == nil {
		t.Fatalf("DeleteList error: want error, got result uid = %d", res.Uid)
	}
	want = fmt.Sprintf("TodoList with uid: %d doesn't empty", 0)
	if err.Error() != want {
		t.Fatalf("DeleteList error: want %s, got %s", want, err.Error())
	}
}
