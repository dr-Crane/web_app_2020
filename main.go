package main

import (
	// "database/sql"
	"fmt"
	"html/template"
	"net/http"
	"os"
	// _ "github.com/go-sql-driver/mysql"
)

func index(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("index.html", "header.html", "footer.html")
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}
	full_text, err := os.ReadFile("full_text.txt")
	var space string = string(full_text)
	t.ExecuteTemplate(w, "index", space)
}

func send(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("form.html", "header.html", "footer.html")
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}
	t.ExecuteTemplate(w, "form", nil)
}

// func save_message(w http.ResponseWriter, r *http.Request) {
// 	message := r.FormValue("message")
// 	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3307)/messages_web_app")
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer db.Close()
// 	fmt.Println("Confirmed!!!")
// 	insert, err := db.Query(fmt.Sprintf("INSERT INFO 'messages' ('text') VALUES('%s')", message))
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer insert.Close()
// 	http.Redirect(w, r, "/", http.StatusSeeOther)
// }

func save_message(w http.ResponseWriter, r *http.Request) {
	message := r.FormValue("message")
	file, err := os.Create("message.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	full_text, err := os.Create("full_text.txt")
	if err != nil {
		panic(err)
	}
	defer full_text.Close()
	var result string
	for i := 0; i < len(message); i++ {
		if message[i] == ' ' {
			if message[i+1] == ' ' {
				result = result + "1"
				i = i + 1
			} else {
				result = result + "0"
			}
		}
	}
	file.WriteString(result)
	full_text.WriteString(message)
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func handleFunc() {
	http.HandleFunc("/", index)
	http.HandleFunc("/send", send)
	http.HandleFunc("/save_message", save_message)
	http.ListenAndServe(":8080", nil)
}

func main() {
	handleFunc()
}
