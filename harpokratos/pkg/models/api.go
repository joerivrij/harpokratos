package models

type ResultModel struct {
	Result string `json:"result"`
}

type SecretResultModel struct {
	Result Secret `json:"result"`
}

type SecretRequest struct {
	EngineName string `json:"engineName"`
	SecretName string `json:"secretName"`
	DataOnly bool `json:"dataOnly"`
}
