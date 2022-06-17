package popup

import (
	"fmt"
	"os"
)

func Popup() string {
	js, err := os.ReadFile("components/popup/popup.js")
	if err != nil {
		fmt.Println(err.Error())
	}
	div := `<div class="popup" id="popup"></div><script>%s</script>`
	return fmt.Sprintf(div, js)
}
