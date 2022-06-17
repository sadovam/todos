package todo

import (
	"fmt"

	"github.com/sadovam/todos/models"
)

func TodoItem(d *models.TodoItem) string {
	button := fmt.Sprintf(`<button class="todo__btn" onclick="layout__del(%d)">x</button>`, d.Uid)
	class := "todo__title"
	if d.IsDone {
		class += " todo__done"
	}
	li := fmt.Sprintf(`<li id="layout %d" class="todo__item"><p class="%s">%s</p> %s</li>`, d.Uid, class, d.Title, button)
	return li
}
