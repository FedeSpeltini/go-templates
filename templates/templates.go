package templates

import (
	"os"
	"text/template"

	structs "github.com/FedeSpeltini/go-template/structs"
)

func TextTemplate() {

	per1 := structs.Person{Name: "Alice", Age: 30, Bands: []string{"Beatles", "Led Zeppelin"}}
	loadTemplate("template.txt", per1)
	//tmp, err := template.New("template.txt").ParseFiles("Templates/template.txt")

}

var funcs = template.FuncMap{
	"increment": func(i int) int {
		return i + 1
	},
}

func HtmlTemplate() {
	per1 := structs.Person{Name: "Alice", Age: 30, Bands: []string{"Beatles", "Led Zeppelin"}}
	loadTemplate("index.html", per1)
}

func loadTemplate(filenName string, data interface{}) {
	tmp, err := template.New(filenName).Funcs(funcs).ParseFiles("templates/" + filenName)
	if err != nil {
		panic(err)
	}
	tmp.Execute(os.Stdout, data)
}
