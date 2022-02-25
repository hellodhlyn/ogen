package main

import (
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/hellodhlyn/ogen/internal"
)

func generate(w http.ResponseWriter, req *http.Request) {
	q := req.URL.Query()
	title := q.Get("title")
	profileImageURL := q.Get("profile_image")
	authorName := q.Get("author")

	w.Header().Set("Content-Type", "image/svg+xml")
	svg := internal.GenerateSVG(req.Context(), &internal.Payload{
		Title:           title,
		ProfileImageURL: profileImageURL,
		AuthorName:      authorName,
	})
	_, _ = io.Copy(w, svg)
}

func main() {
	http.Handle("/generate", http.HandlerFunc(generate))
	fmt.Printf("listening port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
