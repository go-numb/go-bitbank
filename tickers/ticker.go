package ticker

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"path"
	"strings"

	"gitlab.com/k-terashima/go-bitbank/types"
)

const (
	BASEURL = "https://public.bitbank.cc/"
	PATH    = "ticker"
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
	Code      int        `json:"code"`
	Sell      float64    `json:"sell,string"`
	Buy       float64    `json:"buy,string"`
	High      float64    `json:"high,string"`
	Low       float64    `json:"low,string"`
	Last      float64    `json:"last,string"`
	Vol       float64    `json:"vol,string"`
	Timestamp types.Time `json:"timestamp"`
}

// Set pair string
func (p *Request) Set(pair string) {
	p.Pair = strings.ToLower(pair)
}

func (p *Request) Get() (Ticker, error) {
	url := BASEURL + path.Join(p.Pair, PATH)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return Ticker{}, err
	}
	// req.Header.Set("Content-Type", "application/json")

	c := new(http.Client)
	res, err := c.Do(req)
	if err != nil {
		return Ticker{}, err
	}
	defer res.Body.Close()

	var resp Response
	json.NewDecoder(res.Body).Decode(&resp)
	if resp.Success != 1 {
		return Ticker{}, errors.New(fmt.Sprintf("response error, not success. error code is %d", resp.Data.Code))
	}

	return resp.Data, nil
}
