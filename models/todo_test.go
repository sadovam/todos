package models

import "testing"

func TestTodoItemIsEqual(t *testing.T) {
	item1 := TodoItem{Uid: 12, ListUid: 2, Title: "TodoItem12", IsDone: false}
	todos := []TodoItem{
		{Uid: 1, ListUid: 2, Title: "TodoItem12", IsDone: false},
		{Uid: 12, ListUid: 3, Title: "TodoItem12", IsDone: false},
		{Uid: 12, ListUid: 2, Title: "Todo12", IsDone: false},
		{Uid: 12, ListUid: 2, Title: "TodoItem12", IsDone: true},
		{Uid: 12, ListUid: 2, Title: "TodoItem12", IsDone: false},
	}
	want := []bool{true, false, false, false, true}
	for i, item := range todos {
		got := item1.IsEqual(&item)
		if got != want[i] {
			t.Fatalf("TodoItem IsEqual error: item1 %v, item2 %v, want %t, get %t",
				item1, item, want[i], got)
		}
	}
}

func TestTodoItemIsSame(t *testing.T) {
	item1 := TodoItem{Uid: 12, ListUid: 2, Title: "TodoItem12", IsDone: false}
	todos := []TodoItem{
		{Uid: 1, ListUid: 2, Title: "TodoItem12", IsDone: false},
		{Uid: 12, ListUid: 3, Title: "TodoItem12", IsDone: false},
		{Uid: 12, ListUid: 2, Title: "Todo12", IsDone: false},
		{Uid: 12, ListUid: 2, Title: "TodoItem12", IsDone: true},
		{Uid: 12, ListUid: 2, Title: "TodoItem12", IsDone: false},
	}
	want := []bool{false, false, false, false, true}
	for i, item := range todos {
		got := item1.IsSame(&item)
		if got != want[i] {
			t.Fatalf("TodoItem IsSame error: item1 %v, item2 %v, want %t, get %t",
				item1, item, want[i], got)
		}
	}
}

func TestTodoListIsEqual(t *testing.T) {

	todos1 := []TodoItem{
		{Uid: 12, ListUid: 2, Title: "TodoItem12", IsDone: false},
		{Uid: 1, ListUid: 2, Title: "TodoItem12", IsDone: false},
		{Uid: 12, ListUid: 3, Title: "TodoItem12", IsDone: false},
	}

	todos2 := []TodoItem{
		{Uid: 12, ListUid: 2, Title: "TodoItem12", IsDone: false},
		{Uid: 12, ListUid: 2, Title: "Todo12", IsDone: false},
		{Uid: 12, ListUid: 2, Title: "TodoItem12", IsDone: true},
	}

	todos3 := []TodoItem{
		{Uid: 12, ListUid: 2, Title: "TodoItem12", IsDone: false},
		{Uid: 12, ListUid: 2, Title: "Todo12", IsDone: false},
	}

	list1 := TodoList{Uid: 5, Title: "List 5", Todos: todos1}

	lists := []TodoList{
		{Uid: 7, Title: "List 5", Todos: todos1},
		{Uid: 5, Title: "List", Todos: todos1},
		{Uid: 5, Title: "List 5", Todos: todos2},
		{Uid: 5, Title: "List 5", Todos: todos3},
		{Uid: 5, Title: "List 5", Todos: todos1},
	}

	want := []bool{true, false, false, false, true}
	for i, list := range lists {
		got := list1.IsEqual(&list)
		if got != want[i] {
			t.Fatalf("TodoList IsEqual error: list1 %v, list2 %v, want %t, get %t",
				list1, list, want[i], got)
		}
	}
}

func TestTodoListIsSame(t *testing.T) {
	todos1 := []TodoItem{
		{Uid: 12, ListUid: 2, Title: "TodoItem12", IsDone: false},
		{Uid: 1, ListUid: 2, Title: "TodoItem12", IsDone: false},
		{Uid: 12, ListUid: 3, Title: "TodoItem12", IsDone: false},
	}

	todos2 := []TodoItem{
		{Uid: 12, ListUid: 2, Title: "TodoItem12", IsDone: false},
		{Uid: 12, ListUid: 2, Title: "Todo12", IsDone: false},
		{Uid: 12, ListUid: 2, Title: "TodoItem12", IsDone: true},
	}

	todos3 := []TodoItem{
		{Uid: 12, ListUid: 2, Title: "TodoItem12", IsDone: false},
		{Uid: 12, ListUid: 2, Title: "Todo12", IsDone: false},
	}

	list1 := TodoList{Uid: 5, Title: "List 5", Todos: todos1}

	lists := []TodoList{
		{Uid: 7, Title: "List 5", Todos: todos1},
		{Uid: 5, Title: "List", Todos: todos1},
		{Uid: 5, Title: "List 5", Todos: todos2},
		{Uid: 5, Title: "List 5", Todos: todos3},
		{Uid: 5, Title: "List 5", Todos: todos1},
	}

	want := []bool{false, false, false, false, true}
	for i, list := range lists {
		got := list1.IsSame(&list)
		if got != want[i] {
			t.Fatalf("TodoList IsEqual error: list1 %v, list2 %v, want %t, get %t",
				list1, list, want[i], got)
		}
	}
}
