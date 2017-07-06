package main

import (
	"log"
	"net/http"

	"fmt"

	"./config"
	"./routes"
)

func main() {
	configuration := config.LoadConfig()

	routerChan := make(chan routes.RouterRequest, 500)
	go routes.Router(routerChan, configuration.ReglasRuteo)

	fmt.Println(configuration)

	http.HandleFunc("/", handleIndex)

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handleIndex(res http.ResponseWriter, req *http.Request) {
	fmt.Println("Request recibido")
}
