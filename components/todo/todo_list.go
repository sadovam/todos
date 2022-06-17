package todo

import (
	"fmt"
	"os"

	"github.com/sadovam/todos/models"
)

func TodoList(d *models.TodoList) string {

	header := fmt.Sprintf(`<h1 class="todo__header">%s</h1>`, d.Title)
	todos := ""
	for _, todo := range d.Todos {
		todos += TodoItem(todo)
	}
	list := fmt.Sprintf(`<ul id="layout__todos">%s</ul>`, todos)

	form := `<form id="layout__form">
				<label>Title:</label>
				<input id="layout__input" type="text" name="title">
				<button class="todo__btn" onclick="layout__on_click(event)">Add</button>
			</form>`
	js, err := os.ReadFile("components/todo/todo_list.js")
	if err != nil {
		fmt.Println(err.Error())
	}
	block := fmt.Sprintf(`<div class="todos">%s%s%s</div>`, header, list, form)

	script := fmt.Sprintf(`<script>%s</script>`, js)

	return block + script
}
