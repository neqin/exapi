package futures

import (
	"crypto/hmac"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/base64"
	"fmt"
	"net/url"
)

func createSignature(urlPath string, values url.Values, nonce int64, secret []byte) string {
	message := values.Encode() + fmt.Sprint(nonce) + urlPath
	shaSum := getSha256([]byte(message))
	macSum := getHMacSha512(shaSum, secret)
	return base64.StdEncoding.EncodeToString(macSum)
}

func getSha256(input []byte) []byte {
	sha := sha256.New()
	sha.Write(input)
	return sha.Sum(nil)
}

func getHMacSha512(message, secret []byte) []byte {
	mac := hmac.New(sha512.New, secret)
	mac.Write(message)
	return mac.Sum(nil)
}
