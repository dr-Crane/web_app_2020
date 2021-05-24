package main

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
)

func index(w http.ResponseWriter, r *http.Request) {
	tmp, err := template.ParseFiles("templates/index.html")

	if err != nil {
		fmt.Fprint(w, err.Error())
	}

	if err != nil {
		fmt.Fprint(w, err.Error())
	}
	tmp.Execute(w, nil)

}

func send(w http.ResponseWriter, r *http.Request) {
	tmp, err := template.ParseFiles("templates/form.html")

	if err != nil {
		fmt.Fprint(w, err.Error())
	}

	if err != nil {
		fmt.Fprint(w, err.Error())
	}

	tmp.Execute(w, nil)
}

func datasave(w http.ResponseWriter, r *http.Request) {
	message := r.FormValue("message")
	existfile, err := os.Open("message.txt")
	if err != nil {
		fmt.Fprint(w, err.Error())
	}
	existfile.Write([]byte(message))
	existfile.Close()

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func handlerequest() {
	http.HandleFunc("/", index)
	http.HandleFunc("/send", send)
	http.HandleFunc("/datasave", datasave)
	http.ListenAndServe(":8080", nil)
}

func main() {
	handlerequest()
}
