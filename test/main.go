package main

import (
	"log"
	"net/http"
	"io"
	"math/rand"
	"strconv"
	"os"
)

func main() {
	http.HandleFunc("/foo", handleIndex)

	log.Fatal(http.ListenAndServe(":" + os.Args[1], nil))
}

func handleIndex(res http.ResponseWriter, req *http.Request) {
	total := req.URL.Query().Get("total")

	int, _ := strconv.Atoi(total)

	io.WriteString(res, randStringBytes(int) + "\n")
}

func randStringBytes(n int) string {
	letterBytes := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}