package main

import (
	"log"
	"net/http"
	"html/template"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var tpl *template.Template

func init()  {
	tpl = template.Must(template.ParseGlob("templates/*"))

}

type person struct {
	FirstName string
	LastName string
}

func main()  {
	http.HandleFunc("/", foo)
	http.ListenAndServe(":8080", nil)
}



func foo(w http.ResponseWriter, req *http.Request)  {
	f := req.FormValue("first")
	l := req.FormValue("last")
	err := tpl.ExecuteTemplate(w, "index.html", person{f, l})
	if err != nil{
		http.Error(w, err.Error(), 500)
		log.Fatalln(err)
	}
}
