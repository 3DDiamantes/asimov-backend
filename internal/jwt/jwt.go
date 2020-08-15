package jwt

import (
	"asimov-backend/internal/base64"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/json"
)

type Header struct {
	Algorithm string `json:"alg"`
	Type      string `json:"typ"`
}

type Payload struct {
	Subject string `json:"sub"`
}

func GenerateToken(header Header, payload Payload) string {
	headerBytes, _ := json.Marshal(header)
	headerEncoded := base64.Encode(headerBytes)

	payloadBytes, _ := json.Marshal(payload)
	payloadEncoded := base64.Encode(payloadBytes)

	signature := generateSignature(headerEncoded, payloadEncoded)

	return headerEncoded + "." + payloadEncoded + "." + signature
}

func generateSignature(header string, payload string) string {
	secret := base64.EncodeString("secret")

	h := hmac.New(sha256.New, []byte(secret))

	h.Write([]byte(header + "." + payload))

	signature := base64.Encode(h.Sum(nil))
	return signature
}
