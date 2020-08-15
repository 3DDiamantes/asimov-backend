package utils

import "encoding/base64"

func Base64Encode(data string) string {
	return base64.URLEncoding.EncodeToString([]byte(data))
}

func Base64Decode(encoded string) string {
	data, _ := base64.URLEncoding.DecodeString(encoded)
	return string(data)
}
