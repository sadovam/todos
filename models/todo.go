package models

type TodoItem struct {
	Uid     int
	ListUid int
	Title   string
	IsDone  bool
}

func (item1 *TodoItem) IsEqual(item2 *TodoItem) bool {
	return item1.ListUid == item2.ListUid &&
		item1.Title == item2.Title &&
		item1.IsDone == item2.IsDone
}

func (item1 *TodoItem) IsSame(item2 *TodoItem) bool {
	if item1.Uid != item2.Uid {
		return false
	}
	return item1.IsEqual(item2)
}

type TodoList struct {
	Uid   int
	Title string
	Todos []TodoItem
}

func (list1 *TodoList) IsEqual(list2 *TodoList) bool {
	if list1.Title != list2.Title || len(list1.Todos) != len(list2.Todos) {
		return false
	}

	for i, item := range list1.Todos {
		if !item.IsEqual(&list2.Todos[i]) {
			return false
		}
	}

	return true
}

func (list1 *TodoList) IsSame(list2 *TodoList) bool {
	if list1.Uid != list2.Uid {
		return false
	}

	return list1.IsEqual(list2)
}
