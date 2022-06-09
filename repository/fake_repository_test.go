package repository

import "testing"

func TestMakeFakeData(t *testing.T) {
	data := MakeFakeData(5, 7)
	if len(data) != 5 {
		t.Fatalf("Generating fake data error; may create 5 lists, but creating %d", len(data))
	}
	if len(data[0].Todos) != 7 {
		t.Fatalf("Generating fake data error; may create 7 todos per list, but creating %d", len(data))
	}
}

func TestGetItemByUid(t *testing.T) {
	data := MakeFakeData(5, 7)
	repo := NewTodosFakeRepository(data)
	res, err := repo.GetItemByUid(3)

	if err != nil {
		t.Fatalf("GetTestItemByUid error: %v", err.Error())
	}

	if res.Uid != 3 || res.ListUid != 0 {
		t.Fatalf("GetTestItemByUid error: want uid = 3, list uid = 0, got uid = %d, list uid = %d", res.Uid, res.ListUid)
	}
}

func TestGetItemByUidError(t *testing.T) {
	data := MakeFakeData(5, 7)
	repo := NewTodosFakeRepository(data)
	res, err := repo.GetItemByUid(1333)

	if err == nil {
		t.Fatalf("GetTestItemByUid error: want error, got result uid = %d, list uid = %d", res.Uid, res.ListUid)
	}
	want := "Can't find todo item with id: 1333"
	if err.Error() != want {
		t.Fatalf("GetTestItemByUid error: want %s, got %s", want, err.Error())
	}
}

func TestGetItemsByListUid(t *testing.T) {
	data := MakeFakeData(5, 7)
	repo := NewTodosFakeRepository(data)
	res, err := repo.GetItemsByListUid(3)

	if err != nil {
		t.Fatalf("GetTestItemsByListUid error: %v", err.Error())
	}

	if res[2].Uid != 23 || res[2].ListUid != 3 {
		t.Fatalf("GetTestItemsByListUid error: want uid = 23, list uid = 3, got uid = %d, list uid = %d", res[2].Uid, res[2].ListUid)
	}
}

func TestGetItemsByListUidError(t *testing.T) {
	data := MakeFakeData(5, 7)
	repo := NewTodosFakeRepository(data)
	res, err := repo.GetItemsByListUid(13)

	if err == nil {
		t.Fatalf("GetTestItemByUid error: want error, got result %v", res)
	}
	want := "Can't find todo items with list uid: 13"
	if err.Error() != want {
		t.Fatalf("GetTestItemByUid error: want %s, got %s", want, err.Error())
	}
}
