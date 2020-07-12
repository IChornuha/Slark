package books

import (
	"fmt"
	"os"
	"strings"

	"github.com/bmaupin/go-epub"
	strip "github.com/grokify/html-strip-tags-go"
)

//BookDirectoryPath contain book files catalog
const BookDirectoryPath = "./files"

//Book structure represent forum topic as json-encoded data
type Book struct {
	Title    string
	Genre    string
	Pairing  string
	Size     string
	Status   string
	Rating   string
	Author   string
	Source   string
	Document *Document
	file     *epub.Epub
}

//Document reprecent book body
type Document struct {
	Chapters []Chapter
}

//Chapter reprecent book body part
type Chapter struct {
	Title   string
	Content string
}

//Prepare book structure using parsed forum posts array
func (book *Book) Prepare(parsedPosts []string) {
	book.file = epub.NewEpub(book.Title)
	book.Document = &Document{}
	book.Document.BuildJSONStruct(parsedPosts)
	if _, err := os.Stat(BookDirectoryPath); os.IsNotExist(err) {
		os.Mkdir(BookDirectoryPath, 0777)
	}
}

//Write book as *.epub file
func (book *Book) Write() {
	book.file.SetAuthor("Some author")
	for _, chapter := range book.Document.Chapters {
		book.file.AddSection(chapter.Content, chapter.Title, "", "")
	}
	err := book.file.Write(fmt.Sprintf("%s/%s.epub", BookDirectoryPath, book.Title))
	if err != nil {
		fmt.Println("Error while writing file ", BookDirectoryPath, "/", book.Title, ".epub")
	} else {
		fmt.Println("Writed file ", BookDirectoryPath, "/", book.Title, ".epub")
	}
}

//BuildJSONStruct build book
func (doc *Document) BuildJSONStruct(parsedPosts []string) {
	for _, text := range parsedPosts {
		chapterTitle := text[:500]
		chapterTitle = strip.StripTags(chapterTitle)
		chapterTitle = fmt.Sprintf("%s", chapterTitle[:61])
		chapterTitle = strings.TrimLeft(chapterTitle, " ")

		chapterText := text
		chapter := Chapter{Title: chapterTitle, Content: chapterText}
		doc.Chapters = append(doc.Chapters, chapter)
	}
}
