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

func main() {
	configuration := config.LoadConfig()

	routerChan := make(chan routes.RouterRequest)
	go routes.Router(routerChan, configuration.ReglasRuteo)

	fmt.Println(configuration)

	http.HandleFunc("/", handleIndex)

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handleIndex(res http.ResponseWriter, req *http.Request) {
	var client = &http.Client{
		Timeout: time.Second * 10,
	}

	req, err := http.NewRequest(req.Method, nextServer(req), req.Body)

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

func nextServer(req *http.Request) string {
	return "http://localhost:8081"
}
