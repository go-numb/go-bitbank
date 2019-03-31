package ticker

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/url"
	"path"
	"strings"

	e "github.com/go-numb/go-bitbank/errors"

	"github.com/go-numb/go-bitbank/types"
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
	u, err := url.ParseRequestURI(BASEURL)
	if err != nil {
		return Ticker{}, err
	}
	u.Path = path.Join(p.Pair, PATH)

	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return Ticker{}, err
	}

	c := new(http.Client)
	res, err := c.Do(req)
	if err != nil {
		return Ticker{}, err
	}
	if res.StatusCode != 200 {
		return Ticker{}, errors.New(res.Status)
	}
	defer res.Body.Close()

	var resp Response
	json.NewDecoder(res.Body).Decode(&resp)
	if resp.Success != 1 {
		return Ticker{}, e.Handler(resp.Data.Code, err)
	}

	return resp.Data, nil
}
