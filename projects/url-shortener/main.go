package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Url struct {
	Url string `json:"url"`
}

func main() {
	http.HandleFunc("/hello", hello)
	http.HandleFunc("/url", shortenUrl)
	fmt.Println("Listening on port 8000")
	http.ListenAndServe(":8000", nil)
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello")
}

func shortenUrl(w http.ResponseWriter, r *http.Request) {
	var url Url
	if r.Method == "POST" {
		fmt.Println("method post")
		err := json.NewDecoder(r.Body).Decode(&url)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
		fmt.Println(url.Url)
	}
}
