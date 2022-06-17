package page

import (
	"fmt"

	"github.com/sadovam/todos/components/popup"
)

func Page(content string) string {
	header := `<!DOCTYPE html>
	<html lang="en">
	<head>
	   <meta charset="UTF-8">
	   <meta http-equiv="X-UA-Compatible" content="IE=edge">
	   <meta name="viewport" content="width=device-width, initial-scale=1.0">
	   <link rel="stylesheet" href="static/css/discard.css">
	   <link rel="stylesheet" href="static/css/main.css">
	</head>
	<body>`
	popup := popup.Popup()
	footer := `</body></html>`
	return fmt.Sprintf("%s%s%s%s", header, popup, content, footer)
}
