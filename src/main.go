package main

import (
	"log"
	"net/http"

	"fmt"

	"./config"
	"./routes"
	"io/ioutil"
	"time"
)

var routerChan = make(chan routes.RouterRequest, 500)

func main() {
	configuration := config.LoadConfig()

	go routes.Router(routerChan, configuration.ReglasRuteo)

	fmt.Println(configuration)

	http.HandleFunc("/", handleIndex)

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handleIndex(res http.ResponseWriter, req *http.Request) {
	var client = &http.Client{
		Timeout: time.Second * 10,
	}

	req, err := http.NewRequest(req.Method, nextUrl(req), req.Body)

	log.Println(req)

	if err != nil {
		log.Fatal(err)
	}

	response, err := client.Do(req)

	if err != nil {
		log.Fatal(err)
	}

	bodyBytes, err := ioutil.ReadAll(response.Body)

	if err != nil {
		log.Fatal(err)
	}

	res.Write(bodyBytes)
}

func nextUrl(req *http.Request) string {
	path := "/" + req.URL.Path[1:]
	miChannel := make(chan routes.RouterResponse)

	request := routes.RouterRequest{Operation: routes.GiveMeAServer, Path: path, Method: req.Method, C: &miChannel}

	routerChan <- request

	response := <- miChannel

	return "http://" + response.Server + path
}
