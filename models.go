package bitbank

import (
	"sync"

	depth "github.com/go-numb/go-bitbank/depths"
	"github.com/go-numb/go-bitbank/ohlcv"
	private "github.com/go-numb/go-bitbank/privates"
	ticker "github.com/go-numb/go-bitbank/tickers"
	transaction "github.com/go-numb/go-bitbank/transactions"
)

const (
	PUBLIC_URL  = "https://public.bitbank.cc/"
	PRIVATE_URL = "https://private.bitbank.cc/"
	// API 公式ドキュメントに記載ないため，エラーで判別
	// 1mでリミットを設けておく（参考: BitflyerAPI limit）
	API_PRIVATE_LIMIT = 200
	API_PUBLIC_LIMIT  = 500
)

type Client struct {
	PublicURL  string
	PrivateURL string

	Transactions *transaction.Request
	Depth        *depth.Request
	Ticker       *ticker.Request
	OHLCV        *ohlcv.Request

	Auth *private.Auth

	Limit *APILimit
}

func New(token, secret string) *Client {
	return &Client{
		PublicURL:  PUBLIC_URL,
		PrivateURL: PRIVATE_URL,

		Transactions: &transaction.Request{},
		Depth:        &depth.Request{},
		Ticker:       &ticker.Request{},
		OHLCV:        &ohlcv.Request{},

		Auth: private.New(token, secret),

		Limit: &APILimit{
			Public:  API_PUBLIC_LIMIT,
			Private: API_PRIVATE_LIMIT,
		},
	}
}

type APILimit struct {
	sync.Mutex

	Public  int
	Private int
}

func (p *APILimit) Sub(isPrivate bool) {
	p.Lock()
	defer p.Unlock()

	if isPrivate {
		p.Private--
	}
	p.Public--
}

func (p *APILimit) Reset() {
	p.Lock()
	defer p.Unlock()

	p.Private = API_PRIVATE_LIMIT
	p.Public = API_PUBLIC_LIMIT
}
