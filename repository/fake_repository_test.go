package repository

import (
	"fmt"
	"testing"
)

const (
	totalLists   = 5
	itemsPerList = 7

	itemUid           = 3
	itemUIdOutOfRange = 1333

	listUid           = 3
	listUidOutOfRange = 13
	itemPosition      = 3
	itemUidOnPosition = listUid*itemsPerList + itemPosition
)

func TestMakeFakeData(t *testing.T) {
	data := MakeFakeData(totalLists, itemsPerList)
	if len(data) != totalLists {
		t.Fatalf("Generating fake data error; may create %d lists, but creating %d",
			totalLists, len(data))
	}
	if len(data[0].Todos) != itemsPerList {
		t.Fatalf("Generating fake data error; may create %d todos per list, but creating %d",
			itemsPerList, len(data))
	}
}

func TestGetItemByUid(t *testing.T) {
	data := MakeFakeData(totalLists, itemsPerList)
	repo := NewTodosFakeRepository(data)
	res, err := repo.GetItemByUid(itemUid)

	if err != nil {
		t.Fatalf("GetTestItemByUid error: %v", err.Error())
	}

	if res.Uid != itemUid {
		t.Fatalf("GetTestItemByUid error: want uid = %d, got uid = %d", itemUid, res.Uid)
	}
}

func TestGetItemByUidError(t *testing.T) {
	data := MakeFakeData(totalLists, itemsPerList)
	repo := NewTodosFakeRepository(data)
	res, err := repo.GetItemByUid(1333)

	if err == nil {
		t.Fatalf("GetTestItemByUid error: want error, got result uid = %d", res.Uid)
	}
	want := fmt.Sprintf("Can't find todo item with id: %d", itemUIdOutOfRange)
	if err.Error() != want {
		t.Fatalf("GetTestItemByUid error: want %s, got %s", want, err.Error())
	}
}

func TestGetItemsByListUid(t *testing.T) {
	data := MakeFakeData(totalLists, itemsPerList)
	repo := NewTodosFakeRepository(data)
	res, err := repo.GetItemsByListUid(3)

	if err != nil {
		t.Fatalf("GetTestItemsByListUid error: %v", err.Error())
	}

	if res[itemPosition].Uid != itemUidOnPosition || res[itemPosition].ListUid != listUid {
		t.Fatalf(
			"GetTestItemsByListUid error: want uid = %d, list uid = %d, got uid = %d, list uid = %d",
			itemUidOnPosition, listUid, res[itemPosition].Uid, res[itemPosition].ListUid)
	}
}

func TestGetItemsByListUidError(t *testing.T) {
	data := MakeFakeData(totalLists, itemsPerList)
	repo := NewTodosFakeRepository(data)
	res, err := repo.GetItemsByListUid(listUidOutOfRange)

	if err == nil {
		t.Fatalf("GetTestItemByUid error: want error, got result %v", res)
	}
	want := fmt.Sprintf("Can't find todo items with list uid: %d", listUidOutOfRange)
	if err.Error() != want {
		t.Fatalf("GetTestItemByUid error: want %s, got %s", want, err.Error())
	}
}
