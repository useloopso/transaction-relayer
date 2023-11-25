package types

type Config struct {
	NodeUrl string `mapstructure:"NODE_URL"`
	Key     string `mapstructure:"KEY"`
	Mode    string `mapstructure:"MODE"`
	Port    int    `mapstructure:"PORT"`
}

type ExecutePayload struct {
	Address   string `json:"address"`
	Abi       string `json:"abi"`
	Nonce     uint   `json:"nonce"`
	Signature string `json:"signature"`
}

type ExecuteResponse struct {
	TransactionHash string `json:"transactionHash"`
}

type TransactionForSigning struct {
	ExecutePayload  ExecutePayload
	ResponseChannel chan ExecuteResponse
	ErrorChannel    chan error
}

type SignedTransaction struct {
	SignedTransaction string
	TransactionHash   string
}
