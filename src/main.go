package main

import (
	"log"
	"net/http"
	"./routes"
	"os"
	"encoding/json"
	"fmt"
)

type Regla struct {
	Ruta string
	Servers []string
}

type Configuration struct {
	Puerto int
	EsperaRecuperoServer int
	ReglasRuteo []Regla
}

func main() {
	configuration := loadConfig()

	fmt.Println(configuration)

	http.HandleFunc("/", routes.HandleIndex)

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func loadConfig() (Configuration){
	file, _ := os.Open("./config.json")
	decoder := json.NewDecoder(file)
	configuration := Configuration{}
	err := decoder.Decode(&configuration)
	if err != nil {
		fmt.Println("error:", err)
	}
	return configuration
}