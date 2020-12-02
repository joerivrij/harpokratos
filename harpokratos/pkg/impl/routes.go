package impl

import "net/http"

type Adapter func(http.HandlerFunc) http.HandlerFunc

// InitRoutes to start up a mux router and return the routes
func InitRoutes() *http.ServeMux {
	serveMux := http.NewServeMux()

	//pingpong
	serveMux.HandleFunc("/ping", Adapt(PingPong, ValidateRestMethod("GET"), LogRequestDetails()))
	serveMux.HandleFunc("/api/v1/handle", Adapt(SomeOtherHandler, ValidateRestMethod("POST"), LogRequestDetails()))

	return serveMux
}

// Iterate over adapters and run them one by one
func Adapt(h http.HandlerFunc, adapters ...Adapter) http.HandlerFunc {
	for _, adapter := range adapters {
		h = adapter(h)
	}
	return h
}