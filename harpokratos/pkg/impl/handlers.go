package impl

import (
	"encoding/json"
	"fmt"
	"github.com/kpango/glg"
	"harpokratos/pkg/config"
	"harpokratos/pkg/middleware"
	"harpokratos/pkg/models"
	"net/http"
	"strconv"
)

type HarpokratosHandler struct {
	Config *config.HarpokratosConfig
}

// PingPong pongs the ping
func (h *HarpokratosHandler)PingPong(w http.ResponseWriter, req *http.Request) {
	pingPong := models.ResultModel{Result: "pong"}
	middleware.ResponseWithJson(w, pingPong)
}

func (h *HarpokratosHandler)VaultHealthCheck(w http.ResponseWriter, r *http.Request) {
	glg.Debug("connecting to vault to establish health")
	vaultHealth := vaultHealth(h.Config.VaultService, h.Config.VaultToken)

	if !vaultHealth {
		glg.Errorf("%s : %s", "vault healthy", strconv.FormatBool(vaultHealth))
	} else {
		glg.Infof("%s : %s", "vault healthy", strconv.FormatBool(vaultHealth))
	}

	healthy := models.ResultModel{Result: "vault healthy"}
	middleware.ResponseWithJson(w, healthy)
}

func (h *HarpokratosHandler)SecretHandler(w http.ResponseWriter, r *http.Request) {
	var secretRequest models.SecretRequest
	err := json.NewDecoder(r.Body).Decode(&secretRequest)
	if err != nil {
		middleware.ResponseWithJson(w, err)
		return
	}

	vaultLogin, err := vaultLogin(h.Config.VaultUser, h.Config.VaultPassword, h.Config.VaultService)
	if err != nil {
		middleware.ResponseWithJson(w, err)
		return
	}

	vaultURL := fmt.Sprintf("%s/v1/%s/data/%s?metadata=1", h.Config.VaultService, secretRequest.EngineName, secretRequest.SecretName)
	req, _ := http.NewRequest("GET", vaultURL, nil)
	req.Header.Add("X-Vault-Token", vaultLogin.Auth.ClientToken)

	secret, err := getSecretFromVault(*req)
	if err != nil {
		middleware.ResponseWithJson(w, err)
		return
	}

	if secretRequest.DataOnly {
		middleware.ResponseWithJson(w, secret.Data)
		return
	}

	middleware.ResponseWithJson(w, secret)

	return
}