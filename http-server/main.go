package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", hello)
	http.HandleFunc("/headers", headers)
	fmt.Println("Server listening on port 8000")
	http.ListenAndServe(":8000", nil)

}
func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello")
}
func headers(w http.ResponseWriter, r *http.Request) {
	for name, headers := range r.Header {
		for _, h := range headers {
			fmt.Fprintf(w, "%v: %v\n", name, h)
		}
	}
}
