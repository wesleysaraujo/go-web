package main

import (
	"fmt"
	"net/http"
	"html/template"
)

type Post struct {
	Id int
	Title string
	Body string
}

func main() {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		t := template.Must(template.ParseFiles("views/index.html"))
		if err := t.ExecuteTemplate(writer, "index.html", nil); err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
		}
	})
	fmt.Println(http.ListenAndServe(":8080", nil))
}
