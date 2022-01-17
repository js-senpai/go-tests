package main

import (
	"log"
	"os"
	"text/template"
)

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func executeTemplate(text string, data interface{}) {
	tmpl, err := template.New("test").Parse(text)
	check(err)
	err = tmpl.Execute(os.Stdout, data)
	check(err)
}

func main() {
	executeTemplate("Dot is: {{.}}!\n", "ABC")
	executeTemplate("Dot is: {{.}}!\n", 123.5)
	executeTemplate("start {{if .}}Dot is true!{{end}} finish\n", true)
	executeTemplate("start {{if .}}Dot is true!{{end}} finish\n", false)
	templateText := "Before loop: {{.}}\n{{range .}}In loop: {{.}}\n{{end}}After loop: {{.}}\n"
	executeTemplate(templateText, []string{"do", "re", "mi"})
	templateText = "Prices:\n{{range .}}${{.}}\n{{end}}"
	executeTemplate(templateText, []float64{1.25, 0.99, 27})
	type Part struct {
		Name  string
		Count int
	}
	templateStructText := "Name: {{.Name}}\nCount: {{.Count}}\n"
	executeTemplate(templateStructText, Part{Name: "Fuses", Count: 5})
	executeTemplate(templateStructText, Part{Name: "Cables", Count: 2})
}
