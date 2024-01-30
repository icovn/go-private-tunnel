package main

import (
	"errors"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"regexp"
)

var templates = template.Must(template.ParseFiles("tmpl/edit.html", "tmpl/view.html"))
var validPath = regexp.MustCompile("^/(edit|save|view)/([a-zA-Z0-9]+)$")

type Page struct {
	Title string
	Body  []byte
}

func (page *Page) save() error {
	filename := "data/" + page.Title + ".txt"
	return os.WriteFile(filename, page.Body, 0600)
}

func getTitle(writer http.ResponseWriter, request *http.Request) (string, error) {
	matches := validPath.FindStringSubmatch(request.URL.Path)
	if matches == nil {
		http.NotFound(writer, request)
		return "", errors.New("Invalid Page Title")
	}
	return matches[2], nil // The title is the second subexpression.
}

func loadPage(title string) (*Page, error) {
	filename := "data/" + title + ".txt"
	body, error := os.ReadFile(filename)
	if error != nil {
		return nil, error
	}
	return &Page{Title: title, Body: body}, nil
}

func renderTemplate(writer http.ResponseWriter, file string, page *Page) {
	error := templates.ExecuteTemplate(writer, file, page)
	if error != nil {
		http.Error(writer, error.Error(), http.StatusInternalServerError)
		return
	}
}

func makeHandler(fn func(http.ResponseWriter, *http.Request, string)) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		matches := validPath.FindStringSubmatch(request.URL.Path)
		if matches == nil {
			http.NotFound(writer, request)
			return
		}
		fn(writer, request, matches[2])
	}
}

func viewHandler(writer http.ResponseWriter, request *http.Request, title string) {
	page, error := loadPage(title)
	if error != nil {
		http.Redirect(writer, request, "/edit/"+title, http.StatusFound)
		return
	}
	renderTemplate(writer, "view.html", page)
}

func editHandler(writer http.ResponseWriter, request *http.Request, title string) {
	page, error := loadPage(title)
	if error != nil {
		page = &Page{Title: title}
	}
	renderTemplate(writer, "edit.html", page)
}

func saveHandler(writer http.ResponseWriter, request *http.Request, title string) {
	body := request.FormValue("body")
	page := &Page{Title: title, Body: []byte(body)}
	error := page.save()
	if error != nil {
		http.Error(writer, error.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(writer, request, "/view/"+title, http.StatusFound)
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

func rootHandler(writer http.ResponseWriter, request *http.Request) {
	http.Redirect(writer, request, "/view/FrontPage", http.StatusFound)
}

func main() {
	http.HandleFunc("/view/", makeHandler(viewHandler))
	http.HandleFunc("/edit/", makeHandler(editHandler))
	http.HandleFunc("/save/", makeHandler(saveHandler))
	http.HandleFunc("/", rootHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
