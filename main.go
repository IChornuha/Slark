package main

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/IChornuha/slark/src/books"
)

func main() {

	books := books.Books{}
	books.Init()
	templates := template.Must(template.ParseFiles("templates/index.html"))

	http.Handle("/static/",
		http.StripPrefix("/static/",
			http.FileServer(http.Dir("static"))))

	http.Handle("/files/",
		http.StripPrefix("/files/",
			http.FileServer(http.Dir("files"))))

	http.HandleFunc("/", func(w http.ResponseWriter, request *http.Request) {
		if err := templates.ExecuteTemplate(w, "index.html", books); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})
	fmt.Println(http.ListenAndServe(":8888", nil))
}
