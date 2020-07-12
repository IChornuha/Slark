package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"

	"github.com/IChornuha/Slark/src/books"
	"github.com/IChornuha/Slark/src/config"
	"github.com/IChornuha/Slark/src/files"
	"github.com/IChornuha/Slark/src/forum"
)

func main() {

	booksList := files.Books{}
	booksList.Init()
	templates := template.Must(template.ParseFiles("templates/index.html"))

	http.Handle("/static/",
		http.StripPrefix("/static/",
			http.FileServer(http.Dir("static"))))

	http.Handle("/files/",
		http.StripPrefix("/files/",
			http.FileServer(http.Dir("files"))))

	http.HandleFunc("/", func(w http.ResponseWriter, request *http.Request) {
		if err := templates.ExecuteTemplate(w, "index.html", booksList); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	http.HandleFunc("/search", func(w http.ResponseWriter, request *http.Request) {

		type Search struct {
			TopicID int
		}
		search := Search{TopicID: 0}
		topicID := request.FormValue("topicId")
		if topicID != "" {
			search.TopicID, _ = strconv.Atoi(topicID)
		}
		forum := forum.Forum{}
		SlitherinForum := forum.Init()

		SlitherinForum.Auth(config.App.Auth.Login, config.App.Auth.Password)
		SlitherinForum.GetTopic(search.TopicID)
		fmt.Println("Parsed topic: ", SlitherinForum.TopicTitle)

		book := books.Book{Title: SlitherinForum.TopicTitle}
		book.Prepare(SlitherinForum.GetParsedPosts())
		book.Write()
	})

	fmt.Println(http.ListenAndServe(":8888", nil))
}
