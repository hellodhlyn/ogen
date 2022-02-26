package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.Handle("/generate.svg", http.HandlerFunc(getSVG))
	http.Handle("/generate.png", http.HandlerFunc(getPNG))
	fmt.Println("listening port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
