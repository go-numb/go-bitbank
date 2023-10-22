package ticker

import (
	"net/http"
	"path"

	"github.com/go-numb/go-bitbank/v2/types"
)

type Request struct {
	// btc_jpy, xrp_jpy, ltc_btc, eth_btc, mona_jpy, mona_btc, bcc_jpy, bcc_btc
	Pair string
}

type Response struct {
	Success int    `json:"success"`
	Data    Ticker `json:"data"`
}

type Ticker struct {
	Code uint32  `json:"code"`
	Sell float64 `json:"sell,string"`
	Buy  float64 `json:"buy,string"`
	// Last Executed Price
	Last float64 `json:"last,string"`
	// ↓↓ the highest price in last 24 hours
	High      float64    `json:"high,string"`
	Low       float64    `json:"low,string"`
	Open      float64    `json:"open,string"`
	Vol       float64    `json:"vol,string"`
	Timestamp types.Time `json:"timestamp"`
}

func (req *Request) Endpoint() string {
	return types.ENDPOINTPUBLIC
}

func (req *Request) Path() string {
	return path.Join(req.Pair, "ticker")
}

func (req *Request) Method() string {
	return http.MethodGet
}

func (req *Request) Query() string {
	return ""
}

func (req *Request) Payload() []byte {
	return nil
}
