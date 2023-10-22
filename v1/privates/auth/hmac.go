package auth

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"
)

// MakeHeader , option body allow []byte or string
func MakeHeader(token, secret string, body interface{}, req *http.Request) {
	nonce := fmt.Sprintf("%d", time.Now().UnixNano())

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("ACCESS-KEY", token)
	req.Header.Set("ACCESS-NONCE", nonce)
	signs := signature(secret, req.Method, nonce, req.URL.Path, checkBody(body), req.URL.Query())
	req.Header.Set("ACCESS-SIGNATURE", signs)
}

// Signature for Bitbank
// - GETの場合: 「ACCESS-NONCE、リクエストのパス、クエリパラメータ」 を連結させたもの
// - POSTの場合: 「ACCESS-NONCE、リクエストボディのJson文字列」 を連結させたもの
func signature(key, method, nonce, path, body string, q url.Values) string {
	if strings.ToUpper(method) == "GET" {
		nonce += path
		query := q.Encode()
		if query != "" {
			nonce += "?" + query
		}
		return makeHMAC(key, nonce)
	}

	// POSTの場合
	nonce += body
	fmt.Printf("%s\n", nonce)

	return makeHMAC(key, nonce)
}

func makeHMAC(key, str string) string {
	mac := hmac.New(sha256.New, []byte(key))
	mac.Write([]byte(str))
	return hex.EncodeToString(mac.Sum(nil))
}

func checkBody(b interface{}) string {
	switch v := b.(type) {
	case string:
		return v
	case []byte:
		return string(v)
	}

	return ""
}
