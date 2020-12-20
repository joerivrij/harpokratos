package impl

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type ExampleApiConfig struct {
	SecretSauce string `json:"secretSauce"`
	DatabaseConnection string `json:"databaseConnection"`
	Favorites []string `json:"favorites"`
}

func GetConfigFromSideCar(harpokratosUrl string) *ExampleApiConfig {
	time.Sleep(30 * time.Second)

	url := harpokratosUrl + "/api/v1/secret"
	secretRequest := SecretRequest{
		EngineName: "config",
		SecretName: "golangExample",
		DataOnly:   true,
	}
	body, _ := secretRequest.Marshal()

	fmt.Println("requesting: " + url)

	harpoRequest, _ := http.NewRequest("POST", url, bytes.NewBuffer(body))

	client := &http.Client{}
	resp, err := client.Do(harpoRequest)
	if err != nil {
		fmt.Errorf(err.Error())
		return nil
	}
	defer resp.Body.Close()
	responseBody, _ := ioutil.ReadAll(resp.Body)

	fmt.Println("secret: " + string(responseBody))

	secret, err := UnmarshalSecret(responseBody)

	config := &ExampleApiConfig{
		SecretSauce:   secret.Data.SecretSauce,
		DatabaseConnection: secret.Data.DatabaseConnection,
		Favorites: secret.Data.Favorites,
	}

	return config
}