package impl

import "net/http"

// PingPong pongs the ping
func PingPong(w http.ResponseWriter, req *http.Request) {
	pingPong := ResultModel{Result: "pong"}
	ResponseWithJson(w, pingPong)
}

func SomeOtherHandler(w http.ResponseWriter, r *http.Request) {
	randomDrink := ResultModel{Result: "test"}
	ResponseWithJson(w, randomDrink)
}
