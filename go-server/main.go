package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	// fmt.Println("Hello World")
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)

	// other routes handlers
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/hello", helloHandler)

	err := http.ListenAndServe(":8080", nil)

	fmt.Println("Server started on port 8080")

	if err != nil {
		log.Fatal("ListenAndServe: ", err)
		return
	}
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	// check the pathname for the request is hello
	if r.URL.Path != "/hello" {
		http.NotFound(w, r)
		return
	}
	// check the request method is GET
	if r.Method != "GET" {
		http.Error(w, "Method is not supported.", http.StatusMethodNotAllowed)
		return
	}
	fmt.Fprintf(w, "Hello World")
}

func formHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}
	fmt.Fprintf(w, "Post from website! r.PostFrom = %v, r.Form = %v", r.PostForm, r.Form)
	name := r.FormValue("name")
	email := r.FormValue("email")
	fmt.Fprintf(w, "Name = %s, email = %s, r.Form = %v", name, email, r.Form)
}
