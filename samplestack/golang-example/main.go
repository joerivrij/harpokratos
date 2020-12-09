package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = ":5000"
	}
	fmt.Println("running on port: " + port)

	serveMux := http.NewServeMux()
	serveMux.HandleFunc("/ping", PingPong)

	err := http.ListenAndServe(port, serveMux)
	if err != nil {
		panic(err)
	}
}

// PingPong pongs the ping
func PingPong(w http.ResponseWriter, req *http.Request) {
	w.WriteHeader(200)
	w.Write([]byte("pong"))
}
