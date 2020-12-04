package impl

import (
	"harpokratos/pkg/config"
	"harpokratos/pkg/middleware"
	"net/http"
)

// InitRoutes to start up a mux router and return the routes
func InitRoutes(config config.HarpokratosConfig) *http.ServeMux {
	serveMux := http.NewServeMux()

	harpokratorHandler := HarpokratosHandler{Config: &config}
	//pingpong
	serveMux.HandleFunc("/ping", middleware.Adapt(harpokratorHandler.PingPong, middleware.ValidateRestMethod("GET"), middleware.LogRequestDetails()))
	serveMux.HandleFunc("/api/v1/health", middleware.Adapt(harpokratorHandler.VaultHealthCheck, middleware.ValidateRestMethod("GET"), middleware.LogRequestDetails()))
	serveMux.HandleFunc("/api/v1/secret", middleware.Adapt(harpokratorHandler.SecretHandler, middleware.ValidateRestMethod("POST"), middleware.LogRequestDetails()))

	return serveMux
}