package api

import (
	"begginer/url-shortener/data"
	"encoding/json"
	"fmt"
	"net/http"
)

func Post(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Unsupported Method", http.StatusMethodNotAllowed)
	}
	var url data.Url
	err := json.NewDecoder(r.Body).Decode(&url)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	fmt.Println(data.Shorten())
}
