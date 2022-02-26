package main

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/hellodhlyn/ogen/internal"
)

func getSVG(w http.ResponseWriter, req *http.Request) {
	payload, err := parseQuery(req.URL.Query())
	if err != nil {
		_, _ = w.Write([]byte(err.Error()))
		w.WriteHeader(http.StatusBadRequest)
	}

	w.Header().Set("Content-Type", "image/svg+xml")
	svg := internal.GenerateSVG(req.Context(), payload)
	_, _ = io.Copy(w, svg)
}

func getPNG(w http.ResponseWriter, req *http.Request) {
	payload, err := parseQuery(req.URL.Query())
	if err != nil {
		_, _ = w.Write([]byte(err.Error()))
		w.WriteHeader(http.StatusBadRequest)
	}

	svg := internal.GenerateSVG(req.Context(), payload)
	png, err := internal.ConvertSVGToPNG(req.Context(), svg)
	if err != nil {
		fmt.Printf("failed to convert svg to png: %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		w.Header().Set("Content-Type", "image/png")
		_, _ = io.Copy(w, png)
	}
}

func parseQuery(q url.Values) (*internal.Payload, error) {
	title := q.Get("title")
	profileImageURL := q.Get("profile_image")
	authorName := q.Get("author")
	if title == "" || profileImageURL == "" || authorName == "" {
		return nil, errors.New("some required parameters are empty")
	}

	return &internal.Payload{
		Title:           title,
		ProfileImageURL: profileImageURL,
		AuthorName:      authorName,
	}, nil
}
