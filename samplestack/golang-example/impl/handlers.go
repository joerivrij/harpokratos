package impl

import (
	"encoding/json"
	"net/http"
)

// PingPong pongs the ping
func (e *ExampleApiConfig)PingPong(w http.ResponseWriter, req *http.Request) {
	w.WriteHeader(200)
	w.Write([]byte("pong"))
}

func (e *ExampleApiConfig)ConfigHandler(w http.ResponseWriter, req *http.Request) {
	config := ExampleApiConfig{
		SecretSauce:        e.SecretSauce,
		DatabaseConnection: e.DatabaseConnection,
		Favorites:          e.Favorites,
	}

	ResponseWithJson(w, config, 200)
}

// ResponseWithJson returns formed JSON
func ResponseWithJson(w http.ResponseWriter, payload interface{}, code int) {
	response, _ := json.Marshal(payload)
	resp := string(response)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write([]byte(resp))
}