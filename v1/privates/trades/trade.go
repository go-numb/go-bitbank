package trades

import (
	"encoding/json"
	"net/http"
	"net/url"
	"path"

	e "github.com/go-numb/go-bitbank/v1/errors"

	"github.com/google/go-querystring/query"

	"github.com/go-numb/go-bitbank/v1/privates/auth"
	"github.com/go-numb/go-bitbank/v1/types"
)

const (
	BASEURL = "https://api.bitbank.cc/"
	VERSION = "v1"
	PATH    = "user/spot/trade_history"
)

type Request struct {
	Token  string
	Secret string

	Body string
}

type Setter struct {
	Pair        string  `json:"pair" url:"pair,omitempty"`
	OrderID     float64 `json:"order_id" url:"order_id,omitempty"`
	Count       float64 `json:"count" url:"count,omitempty"`
	Since       int64   `json:"since" url:"since,omitempty"`
	End         int64   `json:"end" url:"end,omitempty"`
	IsSortUpper string  `json:"order" url:"order,omitempty"`
}

type Response struct {
	Success int `json:"success"`
	Data    struct {
		Code   int    `json:"code"`
		Trades Trades `json:"trades"`
	} `json:"data"`
}

type Trades []Trade

type Trade struct {
	TradeID        int        `json:"trade_id"`
	Pair           string     `json:"pair"`
	OrderID        int        `json:"order_id"`
	Side           string     `json:"side"`
	Type           string     `json:"type"`
	Amount         float64    `json:"amount,string"`
	Price          float64    `json:"price,string"`
	MakerTaker     float64    `json:"maker_taker,string"`
	FeeAmountBase  float64    `json:"fee_amount_base,string"`
	FeeAmountQuote float64    `json:"fee_amount_quote,string"`
	ExecutedAt     types.Time `json:"executed_at"`
}

func (p *Request) Set(s *Setter) {
	values, _ := query.Values(s)
	p.Body = values.Encode()
}

func (p *Request) Get() (Trades, error) {
	u, err := url.ParseRequestURI(BASEURL)
	if err != nil {
		return Trades{}, err
	}
	u.Path = path.Join(VERSION, PATH)
	u.RawQuery = p.Body

	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return Trades{}, err
	}

	auth.MakeHeader(p.Token, p.Secret, nil, req)

	c := new(http.Client)
	res, err := c.Do(req)
	if err != nil {
		return Trades{}, err
	}
	defer res.Body.Close()

	var resp Response
	json.NewDecoder(res.Body).Decode(&resp)
	if resp.Success != 1 {
		return Trades{}, e.Handler(resp.Data.Code, err)
	}

	return resp.Data.Trades, nil
}
