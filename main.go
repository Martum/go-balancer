package main

import (
	"log"
	"net/http"
	"./routes"
)

func main() {

	// public views
	http.HandleFunc("/", routes.HandleIndex)

	log.Fatal(http.ListenAndServe(":8080", nil))
}