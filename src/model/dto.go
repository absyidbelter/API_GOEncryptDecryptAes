package model

type RequestPayload struct {
	EncSalt string `json:"enc_salt"`
	Value   string `json:"value"`
	Action  string `json:"action"` // "encrypt" or "decrypt"
}
