package main

import (
	"log"
	"net/http"
	"io"
)

func main() {
	http.HandleFunc("/foo", handleIndex)

	log.Fatal(http.ListenAndServe(":8081", nil))
}

func handleIndex(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "hola mundo\n")
}
