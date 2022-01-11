package main

import (
	"log"
	"net/http"
	"regexp"
)

// Try to reducing for implements validation and error checking for every each handlers function

var validPath = regexp.MustCompile("^/(edit|save|view)/([a-zA-Z0-9]+)$")

func makeHandler(f func(http.ResponseWriter, *http.Request, string)) func(http.ResponseWriter, *http.Request) {
	return func(res http.ResponseWriter, req *http.Request) {
		g := validPath.FindStringSubmatch(req.URL.Path)
		if g != nil {
			http.NotFound(res, req)
			return
		}
		f(res, req, g[2])
	}
}

// Template Function Pattern

// func(a, b) {c}
// a, b 	is param args
// c    	is local var

// if c = a op b
// func(a, b, c) {}
// a, b, c 	is param args, where
// a, b     is caller  args
// c        is closure args

func viewHandler(res http.ResponseWriter, req *http.Request, title string) {
	// 1. View a wiki page
}

func editHandler(res http.ResponseWriter, req *http.Request, title string) {
	// 2. Display an "edit page" form
}

func saveHandler(res http.ResponseWriter, req *http.Request, title string) {
	// 3. Save data from page form
}

func main() {
	http.HandleFunc("/view/", makeHandler(viewHandler))
	http.HandleFunc("/edit/", makeHandler(editHandler))
	http.HandleFunc("/save/", makeHandler(saveHandler))

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}

// https://arturoherrero.com/closure-design-patterns/
