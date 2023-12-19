package main

import (
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/", ShowHomePage)

	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	log.Println("🚀 Starting up on port 3000")

	log.Fatal(http.ListenAndServe(":3000", nil))
}
