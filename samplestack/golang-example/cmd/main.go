package main

import (
	"fmt"
	"golang-example/impl"
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	harpokratosUrl := os.Getenv("HARPOKRATOS_URL")
	if port == "" {
		port = ":5000"
	}

	if harpokratosUrl == "" {
		harpokratosUrl = "http://0.0.0.0:5001"
	}
	fmt.Println("running on port: " + port)

	exampleApiConfig := impl.GetConfigFromSideCar(harpokratosUrl)

	serveMux := http.NewServeMux()
	serveMux.HandleFunc("/ping", exampleApiConfig.PingPong)
	serveMux.HandleFunc("/api/v1/config", exampleApiConfig.ConfigHandler)

	err := http.ListenAndServe(port, serveMux)
	if err != nil {
		panic(err)
	}
}
