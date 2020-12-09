package impl

import (
	"bytes"
	"encoding/json"
	"fmt"
	vault "github.com/hashicorp/vault/api"
	"github.com/kpango/glg"
	"net/http"
)

func vaultHealth(vaultURL, vaultToken string) bool {
	vaultURL = vaultURL + "/v1/secret/data"
	req, _ := http.NewRequest("GET", vaultURL, nil)
	req.Header.Add("X-Vault-Token", vaultToken)

	_, err := getSecretFromVault(*req)
	if err != nil {
		return false
	}

	glg.Info("connection made to vault")
	return true
}

func getSecretFromVault(vaultRequest http.Request) (*vault.Secret, error) {
	client := &http.Client{}
	resp, err := client.Do(&vaultRequest)
	if err != nil {
		return nil, err
	}

	secret, err := vault.ParseSecret(resp.Body)
	if err != nil {
		return nil, err
	}

	return secret, nil
}

func vaultLogin(username, password, vaultURL string) (*vault.Secret, error) {
	url := fmt.Sprintf("%s/v1/auth/userpass/login/%s", vaultURL, username)
	requestBody, err := json.Marshal(map[string]string{
		"password": password,
	})
	if err != nil {
		return nil, err
	}
	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(requestBody))
	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	secret, err := vault.ParseSecret(resp.Body)
	if err != nil {
		return nil, err
	}

	return secret, nil
}