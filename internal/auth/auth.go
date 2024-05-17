package auth

import (
	"errors"
	"net/http"
	"strings"
)

// GetAPIKey extracts the API key from the request headers
// and returns it. If the API key is not present, it returns
// Example: Authorization: ApiKey <key>
func GetAPIKey(headers http.Header) (string, error) {
	val := headers.Get("Authorization")
	if val == "" {
		return "", errors.New("no authentication header found")
	}

	vals := strings.Split(val, " ")
	if len(vals) != 2 {
		return "", errors.New("invalid authentication header")
	}

	if vals[0] != "ApiKey" {
		return "", errors.New("malformed authentication header prefix")
	}

	return vals[1], nil
}
