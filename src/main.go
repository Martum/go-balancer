package main

import (
	"log"
	"net/http"

	"./config"
	"./routes"
	"io/ioutil"
	"time"
	"strconv"
)

var routerChan = make(chan routes.RouterRequest, 500)

func main() {
	configuration := config.LoadConfig()

	go routes.Router(routerChan, configuration.ReglasRuteo, configuration.EsperaRecuperoServer)

	http.HandleFunc("/", handleIndex)

	log.Fatal(http.ListenAndServe(":" + strconv.Itoa(configuration.Puerto), nil))
}

func handleIndex(res http.ResponseWriter, req *http.Request) {
	response, statusError := forwardRequest(req)

	if response == nil {
		res.WriteHeader(statusError)
	} else {
		res.WriteHeader(response.StatusCode)

		bodyBytes, _ := ioutil.ReadAll(response.Body)

		res.Write(bodyBytes)
	}
}

func forwardRequest(req *http.Request) (*http.Response, int){
	var client = &http.Client{
		Timeout: time.Second * 5,
	}

	for {
		server := nextServer(req)

		if server.RouteRequest {
			request, _ := http.NewRequest(req.Method, makeUrl(server.Server, req.URL.Path), req.Body)
			response, _ := client.Do(request)

			if response.StatusCode == http.StatusRequestTimeout {
				notifyServerDown(request.URL.Path, server.Server)

				if req.Method != http.MethodGet {
					return nil, http.StatusServiceUnavailable;
				}
			} else {
				return response, http.StatusOK
			}
		} else {
			return nil, http.StatusServiceUnavailable
		}
	}
}

func nextServer(req *http.Request) routes.RouterResponse {
	miChannel := make(chan routes.RouterResponse)

	request := routes.RouterRequest{Operation: routes.GiveMeAServer, Path: req.URL.Path, C: &miChannel}

	routerChan <- request

	response := <- miChannel

	return response
}

func makeUrl(server string, path string) string {
	return "http://" + server + path
}

func notifyServerDown(path string, server string) {
	request := routes.RouterRequest{Operation: routes.ServerDown, Path: path, C: nil, Meta: server}
	routerChan <- request
}
