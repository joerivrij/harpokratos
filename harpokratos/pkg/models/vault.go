package models

import "encoding/json"

func UnmarshalVaultSecret(data []byte) (Secret, error) {
	var r Secret
	err := json.Unmarshal(data, &r)
	return r, err
}

type Secret struct {
	RequestID       string          `json:"request_id"`
	LeaseID         string          `json:"lease_id"`
	Renewable       bool            `json:"renewable"`
	LeaseDuration   int64           `json:"lease_duration"`
	VaultSecretData VaultSecretData `json:"data"`
	WrapInfo        interface{}     `json:"wrap_info"`
	Warnings        interface{}     `json:"warnings"`
	Auth            interface{}     `json:"auth"`
}
type VaultSecretData struct {
	Data     interface{} `json:"data"`
	Metadata Metadata    `json:"metadata"`
}

type Metadata struct {
	CreatedTime  string `json:"created_time"`
	DeletionTime string `json:"deletion_time"`
	Destroyed    bool   `json:"destroyed"`
	Version      int64  `json:"version"`
}