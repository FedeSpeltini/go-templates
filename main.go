package main

import (
	"github.com/FedeSpeltini/go-template/commands"
	templates "github.com/FedeSpeltini/go-template/templates"
)

func main() {
	input, err := commands.GetInput()

	if err != nil {
		panic(err)
	}

	switch input {
	case "1":
		templates.TextTemplate()
		break
	case "2":
		templates.HtmlTemplate()
		break

	}
}
