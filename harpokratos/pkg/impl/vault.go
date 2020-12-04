package impl

import (
	"github.com/kpango/glg"
	"harpokratos/pkg/models"
	"io/ioutil"
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

func getSecretFromVault(vaultRequest http.Request) (*models.Secret, error) {
	client := &http.Client{}
	resp, err := client.Do(&vaultRequest)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	responseBody, _ := ioutil.ReadAll(resp.Body)
	glg.Info(string(responseBody))

	vaultSecret, err := models.UnmarshalVaultSecret(responseBody)
	if err != nil {
		return nil, err
	}
	return &vaultSecret, nil
}
