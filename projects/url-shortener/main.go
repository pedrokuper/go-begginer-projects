package main

import (
	"fmt"
	"net/http"
)

func main() {
	server := http.NewServeMux()
	server.HandleFunc("/url", handleRoot)

	fs := http.FileServer(http.Dir("./public"))
	server.Handle("/", fs)

	fmt.Println("Server listening on port 8000")
	err := http.ListenAndServe(":8000", server)
	if err != nil {
		fmt.Println("Error opening the server")
	}
}

func handleRoot(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/url" {
		// fmt.Fprintf(w, "url shortener!")
		w.Write([]byte("url shortener!"))
	}
}
