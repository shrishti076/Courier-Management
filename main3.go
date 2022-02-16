package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

type hotel struct {
	Name    string
	Address string
	city    string
	zip     string
	region  string
}
type hotels []hotel

func init() {
	tpl = template.Must(template.ParseFiles("index2.gohtml"))
}

func main() {
	h := hotels{
		hotel{
			Name:    "Broadway",
			Address: "India",
			city:    "Maharastra",
			zip:     "562101",
			region:  "southern",
		},
		hotel{
			Name:    "Cadway",
			Address: "America",
			city:    "New York",
			zip:     "705401",
			region:  "nothern",
		},
	}
	err := tpl.Execute(os.Stdout, h)
	if err != nil {
		log.Fatalln(err)
	}
}
