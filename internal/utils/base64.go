package utils

import "encoding/base64"

func Base64Encode(data string) string {
	return base64.StdEncoding.EncodeToString([]byte(data))
}

func Base64Decode(encoded string) string {
	data, _ := base64.StdEncoding.DecodeString(encoded)
	return string(data)
}
