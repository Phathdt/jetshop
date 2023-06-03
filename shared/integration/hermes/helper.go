package hermes

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"sort"
	"strings"
)

type hermesHelper struct {
	clientId     string
	clientSecret string
}

func NewHermesHelper(clientId string, clientSecret string) *hermesHelper {
	return &hermesHelper{clientId: clientId, clientSecret: clientSecret}
}

func (h hermesHelper) Sign(params map[string]string, path string, appSecret string) string {
	keys := make([]string, 0)
	for key := range params {
		if key != "sign" {
			keys = append(keys, key)
		}
	}

	sort.StringSlice(keys).Sort()

	payload := path
	for _, key := range keys {
		if params[key] != "" {
			payload = payload + key + params[key]
		}
	}

	return hmacSHA256(payload, appSecret)
}

func hmacSHA256(payload string, appSecret string) string {
	secret := []byte(appSecret)
	bb := []byte(payload)
	hash := hmac.New(sha256.New, secret)
	hash.Write(bb)

	sum := hash.Sum(nil)
	return strings.ToUpper(hex.EncodeToString(sum))
}
