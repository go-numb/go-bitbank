package rest

import (
	"net/http"
	"time"

	"github.com/go-numb/go-bitbank/v2/auth"
	jsoniter "github.com/json-iterator/go"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

type Client struct {
	Auth *auth.Config

	HTTPC *http.Client
}

func New(auth *auth.Config) *Client {
	hc := http.DefaultClient
	hc.Timeout = 5 * time.Second

	return &Client{
		Auth:  auth,
		HTTPC: hc,
	}
}
