package impl

import "encoding/json"

type ResultModel struct {
	Result string `json:"result"`
}

type SecretRequest struct {
	EngineName string `json:"engineName"`
	SecretName string `json:"secretName"`
	DataOnly bool `json:"dataOnly"`
}

func (r *SecretRequest) Marshal() ([]byte, error) {
	return json.Marshal(r)
}


type Config struct {
	SecretSauce string `json:"secretSauce"`
	DatabaseConnection string `json:"databaseConnection"`
	Favorites []string `json:"favorites"`
}


func UnmarshallConfig(data []byte) (Config, error) {
	var r Config
	err := json.Unmarshal(data, &r)
	return r, err
}

func UnmarshalSecret(data []byte) (Secret, error) {
	var r Secret
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *Secret) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type Secret struct {
	Data     Config   `json:"data"`
	Metadata Metadata `json:"metadata"`
}

type Metadata struct {
	CreatedTime  string `json:"created_time"`
	DeletionTime string `json:"deletion_time"`
	Destroyed    bool   `json:"destroyed"`
	Version      int64  `json:"version"`
}