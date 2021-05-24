package main

import (
	"fmt"
	"html/template"
	"net/http"
)

// transit ..
type transit struct {
	message string
}

func index(w http.ResponseWriter, r *http.Request) {
	var bob transit
	tmp, err := template.ParseFiles("templates/index.html")

	if err != nil {
		fmt.Fprint(w, err.Error())
	}

	if err != nil {
		fmt.Fprint(w, err.Error())
	}
	tmp.Execute(w, bob)

}

func send(w http.ResponseWriter, r *http.Request) {
	var bob transit
	tmp, err := template.ParseFiles("templates/form.html")

	if err != nil {
		fmt.Fprint(w, err.Error())
	}

	if err != nil {
		fmt.Fprint(w, err.Error())
	}

	tmp.Execute(w, bob)
}

func handlerequest() {
	http.HandleFunc("/", index)
	http.HandleFunc("/send", send)
	http.ListenAndServe(":8080", nil)
}

func main() {
	handlerequest()
}
