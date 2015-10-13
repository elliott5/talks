// Copyright 2010 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"
	"strings"
)

type Page struct {
	Title string
	Body  []byte
}

func (p *Page) save() error {
	filename := p.Title + ".txt"
	return ioutil.WriteFile(filename, p.Body, 0600)
}

func loadPage(title string) (*Page, error) {
	filename := title + ".txt"
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}

func viewHandler(w http.ResponseWriter, r *http.Request, title string) {
	p, err := loadPage(title)
	if err != nil {
		http.Redirect(w, r, "/edit/"+title, http.StatusFound)
		return
	}
	renderTemplate(w, "view", p)
}

func editHandler(w http.ResponseWriter, r *http.Request, title string) {
	p, err := loadPage(title)
	if err != nil {
		p = &Page{Title: title}
	}
	renderTemplate(w, "edit", p)
}

func saveHandler(w http.ResponseWriter, r *http.Request, title string) {
	body := r.FormValue("body")
	p := &Page{Title: title, Body: []byte(body)}
	err := p.save()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/view/"+title, http.StatusFound)
}

var templates = template.Must(template.ParseFiles("edit.html", "view.html"))

func renderTemplate(w http.ResponseWriter, tmpl string, p *Page) {
	err := templates.ExecuteTemplate(w, tmpl+".html", p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

var validPath = regexp.MustCompile("^/(edit|save|view)/([a-zA-Z0-9]+)$")

func makeHandler(fn func(http.ResponseWriter, *http.Request, string)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		m := validPath.FindStringSubmatch(r.URL.Path)
		if m == nil {
			http.NotFound(w, r)
			return
		}
		fn(w, r, m[2])
	}
}

func main() {
	http.HandleFunc("/view/", makeHandler(viewHandler))
	http.HandleFunc("/edit/", makeHandler(editHandler))
	http.HandleFunc("/save/", makeHandler(saveHandler))
	http.HandleFunc("/new", newTitle)
	http.HandleFunc("/", listTxtFiles)

	http.ListenAndServe(":8080", nil)
}

func listTxtFiles(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "<h1>List of wiki entries</h1>")
	fd, err := os.Open(".")
	if err != nil {
		fmt.Fprintln(w, err)
		return
	}
	fis, err := fd.Readdir(-1)
	if err != nil {
		fmt.Fprintln(w, err)
		return
	}
	for _, fi := range fis {
		if strings.HasSuffix(fi.Name(), ".txt") {
			nam := strings.TrimSuffix(fi.Name(), ".txt")
			fmt.Fprintf(w, `<li><a href="/view/%s">%s</a></li>`, nam, nam)
		}
	}
	fmt.Fprintf(w, `
<form action="/new" method="POST">
<div>New Title:<input type="text" name="title">
<input type="submit" value="New"></div>
</form>
`)
}

func newTitle(w http.ResponseWriter, r *http.Request) {
	title := r.FormValue("title")
	os.Create(title + ".txt")
	http.Redirect(w, r, "/edit/"+title, http.StatusFound)
}
