// package main

// import (
// 	"log"
// 	"os"
// 	"text/template"
// )

// var tpl *template.Template

// type food struct {
// 	Name string
// 	Num  int
// }

// func init() {
// 	tpl = template.Must(template.ParseFiles("index1.gohtml"))
// }

// func main() {

// chocolates := map[string]string{
// 	"first": "shri",
// 	"last":  "shre",
// }

//chocolates := []string{"ab", "bc", "cd"}

// err := tpl.Execute(os.Stdout, chocolates)
// if err != nil {
// 	log.Fatalln(err)
// }

//for maps and slices you can use the under given format to use in the gohtml file or text file
// {{range $key, $value:=.}}
//     <li>{{$key}} {{$value}}</li>
//     {{end}}

//for structs (here the structure is food)
// 	choco := food{
// 		Name: "ferrero rocher",
// 		Num:  5,
// 	}

// 	err := tpl.Execute(os.Stdout, choco)
// 	if err != nil {
// 		log.Fatalln(err)
// 	}

// }