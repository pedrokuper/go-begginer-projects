package shortenurl

import (
	"fmt"
	"math/rand/v2"
)

func ShortenURL(url string) string {
	s := ""
	for i := 0; i < 6; i++ {
		s += string(rand.IntN(26) + 97)
	}

	shortendURL := fmt.Sprintf("http://localhost:8080/%s", s)
	return shortendURL
}
