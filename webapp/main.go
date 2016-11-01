package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
)

var templates = template.Must(template.ParseFiles("templates/edit.html", "templates/view.html"))

type Page struct {
	Title   string
	Content []byte
}

func (p *Page) save() error {
	file_name := p.Title + ".txt"
	return ioutil.WriteFile(file_name, p.Content, 0600)
}

func loadPage(title string) (*Page, error) {
	file_name := title + ".txt"
	content, err := ioutil.ReadFile(file_name)
	if err == nil {
		return &Page{Title: title, Content: content}, nil
	}

	return nil, err
}

func main() {
	// setupPage()
	http.HandleFunc("/view/", wiki)
	http.HandleFunc("/edit/", wikiEdit)
	http.HandleFunc("/save/", wikiSave)
	http.ListenAndServe(":80", nil)

}

func setupPage() {
	page := &Page{Title: "hello", Content: []byte("Hello, World")}
	page.save()
}

func wiki(writer http.ResponseWriter, req *http.Request) {
	page_title := req.URL.Path[len("/view/"):]

	load, err := loadPage(page_title)

	if err == nil {
		renderTemplate(writer, "view", load)
	}

}

func wikiEdit(writer http.ResponseWriter, req *http.Request) {
	page_title := req.URL.Path[len("/edit/"):]
	load, err := loadPage(page_title)

	if err != nil {
		fmt.Println(err)
		http.Redirect(writer, req, "/view/"+page_title, http.StatusFound)
		return
	}
	renderTemplate(writer, "edit", load)

}

func wikiSave(writer http.ResponseWriter, req *http.Request) {
	title := req.FormValue("title")
	content := req.FormValue("content")
	page := &Page{Title: title, Content: []byte(content)}
	err := page.save()
	if err != nil {
		fmt.Println(err)
	} else {
		http.Redirect(writer, req, "/view/"+title, http.StatusFound)
	}

}

func renderTemplate(writer http.ResponseWriter, _temp string, page *Page) {
	// temp, _ := template.ParseFiles("templates/" + _temp + ".html")
	err := templates.ExecuteTemplate(writer, _temp+".html", page)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
	}
	// temp.Execute(writer, page)
}
