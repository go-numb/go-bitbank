package orders

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/url"
	"path"

	e "github.com/go-numb/go-bitbank/errors"

	"github.com/go-numb/go-bitbank/privates/auth"
	"github.com/go-numb/go-bitbank/types"
)

const (
	BASEURL = "https://api.bitbank.cc/"
	VERSION = "v1"
	PATH    = "user/spot"
)

type Request struct {
	Token  string
	Secret string
}

// Body is json body
type Body struct {
	Pair   string  `json:"pair"`
	Amount float64 `json:"amount,string"`
	Price  float64 `json:"price"`
	Side   string  `json:"side"`
	Type   string  `json:"type"`
}

type Response struct {
	Success int `json:"success"`
	Data    struct {
		Code  int   `json:"code"`
		Order Order `json:"order"`
	} `json:"data"`
}

type Order struct {
	OrderID         int        `json:"order_id"`
	Pair            string     `json:"pair"`
	Side            string     `json:"side"`
	Type            string     `json:"type"`
	StartAmount     float64    `json:"start_amount,string"`
	RemainingAmount float64    `json:"remaining_amount,string"`
	ExecutedAmount  float64    `json:"executed_amount,string"`
	Price           float64    `json:"price,string"`
	AveragePrice    float64    `json:"average_price,string"`
	OrderedAt       types.Time `json:"ordered_at"`
	Status          string     `json:"status"`
}

func (p *Request) Post(b *Body) (Order, error) {
	u, err := url.ParseRequestURI(BASEURL)
	if err != nil {
		return Order{}, err
	}
	u.Path = path.Join(VERSION, PATH, "order")

	j, err := json.Marshal(b)
	if err != nil {
		return Order{}, err
	}

	req, err := http.NewRequest("POST", u.String(), bytes.NewReader(j))
	if err != nil {
		return Order{}, err
	}

	auth.MakeHeader(p.Token, p.Secret, j, req)

	c := new(http.Client)
	res, err := c.Do(req)
	if err != nil {
		return Order{}, err
	}
	defer res.Body.Close()

	var resp Response
	json.NewDecoder(res.Body).Decode(&resp)
	if resp.Success != 1 {
		return Order{}, e.Handler(resp.Data.Code, err)
	}

	return resp.Data.Order, nil
}
