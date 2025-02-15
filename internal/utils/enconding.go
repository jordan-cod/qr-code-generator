package utils

import (
	"encoding/base64"
)

func EncodeToBase64(data []byte) string {
	return base64.StdEncoding.EncodeToString(data)
}
