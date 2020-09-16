package utils

import (
	"encoding/base64"
	"encoding/json"
	"strings"
)

// TokenExtractor , for extract data inside token
func TokenExtractor(encodedToken string) int {
	if encodedToken != "" {
		encodedPayload := strings.Split(encodedToken, ".")
		decodedPayload, _ := base64.StdEncoding.DecodeString(encodedPayload[1])

		payload := make(map[string]int)
		_ = json.Unmarshal(decodedPayload, &payload)

		return payload["user_id"]
	}
	return -1
}
