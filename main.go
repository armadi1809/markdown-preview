package main

import (
	"bytes"
	"fmt"
	"html/template"
	"net/http"

	"github.com/microcosm-cc/bluemonday"
	"github.com/yuin/goldmark"
)

func main() {
	http.HandleFunc("/{$}", func(w http.ResponseWriter, r *http.Request) {
		temp, err := template.ParseFiles("./index.html")
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Unable to render home page"))
			return
		}
		err = temp.Execute(w, nil)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Unable to render home page"))
			return
		}
	})
	http.HandleFunc("/parsemdwn", func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		err := r.ParseForm()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Println(err.Error())
			w.Write([]byte("Unable to parse markdown please try again"))
			return
		}
		sanitizedInput := sanitizeInput(r.FormValue("markdown"))

		w.Write(markdownToHTML(sanitizedInput))
	})
	http.ListenAndServe(":3000", nil)
}

func sanitizeInput(input string) string {
	sanitizer := bluemonday.UGCPolicy()
	html := sanitizer.Sanitize(input)
	return html
}

func markdownToHTML(markdown string) []byte {
	var buf bytes.Buffer
	if err := goldmark.Convert([]byte(markdown), &buf); err != nil {
		panic(err)
	}
	return buf.Bytes()
}
