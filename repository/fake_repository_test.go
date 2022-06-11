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
