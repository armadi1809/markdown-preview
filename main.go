package main

import (
	"fmt"
	"html/template"
	"net/http"
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
		w.Write([]byte(fmt.Sprintf("In progress.... markdown for %s goes here", r.FormValue("markdown"))))
	})
	http.ListenAndServe(":3000", nil)
}
