package main

import (
	"begginer/url-shortener/api"
	"fmt"
	"net/http"
)

type Url struct {
	Url string
}

func main() {
	server := http.NewServeMux()
	server.HandleFunc("/url", handleRoot)

	fmt.Println("Server listening on port 8000")
	err := http.ListenAndServe(":8000", server)
	if err != nil {
		fmt.Println("Error opening the server")
	}
}

func handleRoot(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/url" {
		api.Post(w, r)
		// w.Write([]byte("url shortener!"))
	}
}
