package main

import (
<<<<<<< HEAD
=======
	"encoding/json"
>>>>>>> bf9e5d1dc8f49670b5a24aa890e9e60e7158f041
	"fmt"
	"net/http"
)

<<<<<<< HEAD
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
=======
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
>>>>>>> bf9e5d1dc8f49670b5a24aa890e9e60e7158f041
	}
}
