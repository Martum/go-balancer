package routes

import (
	"io"
	"net/http"
)

func HandleIndex(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "hello, world\n")
}