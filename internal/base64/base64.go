package base64

import "encoding/base64"

func Base64EncodeString(data string) string {
	return base64.URLEncoding.EncodeToString([]byte(data))
}

func Base64DecodeString(encoded string) string {
	data, _ := base64.URLEncoding.DecodeString(encoded)
	return string(data)
}

func Base64Encode(data []byte) string {
	return base64.URLEncoding.EncodeToString(data)
}

func Base64Decode(encoded string) []byte {
	data, _ := base64.URLEncoding.DecodeString(encoded)
	return data
}
