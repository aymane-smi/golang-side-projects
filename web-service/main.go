package main

import (
	"fmt"
	"log"
	"net/http"
)

func handleForm(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "error while parsing form!\n%s", err)
		return
	}
	fmt.Fprintf(w, "email:%s\n", r.FormValue("email"))
	fmt.Fprintf(w, "message:%s\n", r.FormValue("message"))
}

func handleHello(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" || r.Method != "GET" {
		http.Error(w, "404 error", http.StatusNotFound)
		return
	}
	fmt.Fprintln(w, "Hello!")
}

func main() {
	fileServer := http.FileServer(http.Dir("./src"))

	http.Handle("/", fileServer)

	http.HandleFunc("/hello", handleHello)

	http.HandleFunc("/form", handleForm)

	fmt.Println("server listening to PORT 8080")

	if err := http.ListenAndServe(":8080", nil); err == nil {
		log.Fatalln("can't listen to port 8080")
	}
}
