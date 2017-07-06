package config

import (
	"encoding/json"
	"fmt"
	"os"
)

type Regla struct {
	Ruta    string
	Servers []string
}

type Configuration struct {
	Puerto               int
	EsperaRecuperoServer int
	ReglasRuteo          []Regla
}

func LoadConfig() Configuration {
	file, _ := os.Open("./config.json")
	decoder := json.NewDecoder(file)
	configuration := Configuration{}
	err := decoder.Decode(&configuration)
	if err != nil {
		fmt.Println("error:", err)
	}
	return configuration
}
