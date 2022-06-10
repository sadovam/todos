package repository

import (
	"fmt"
	"testing"
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
