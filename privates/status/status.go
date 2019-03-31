package status

import (
	"encoding/json"
	"net/http"
	"net/url"
	"strings"

	e "github.com/go-numb/go-bitbank/errors"

	"github.com/go-numb/go-bitbank/privates/auth"
)

const (
	BASEURL = "https://api.bitbank.cc/"
	VERSION = "v1"
	PATH    = "spot/status"
)

type Request struct {
	Token  string
	Secret string
	// btc_jpy, xrp_jpy, ltc_btc, eth_btc, mona_jpy, mona_btc, bcc_jpy, bcc_btc
	Pairs []string
}
type Response struct {
	Success int `json:"success"`
	Data    struct {
		Code     int      `json:"code"`
		Statuses Statuses `json:"statuses"`
	} `json:"data"`
}

type Statuses []Status

type Status struct {
	Pair      string  `json:"pair"`
	Status    string  `json:"status"`
	MinAmount float64 `json:"min_amount,string"`
}

func (p *Request) Set(pairs ...string) {
	for _, v := range pairs {
		p.Pairs = append(p.Pairs, strings.ToLower(v))
	}
}

func (p *Request) Get() (Statuses, error) {
	u, err := url.ParseRequestURI(BASEURL)
	if err != nil {
		return Statuses{}, err
	}
	u.Path = PATH

	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return Statuses{}, err
	}

	auth.MakeHeader(p.Token, p.Secret, nil, req)

	c := new(http.Client)
	res, err := c.Do(req)
	if err != nil {
		return Statuses{}, err
	}
	defer res.Body.Close()

	var resp Response
	json.NewDecoder(res.Body).Decode(&resp)
	if resp.Success != 1 {
		return Statuses{}, e.Handler(resp.Data.Code, err)
	}

	// gets pairs指定がなければすべて返す
	if len(p.Pairs) == 0 {
		return resp.Data.Statuses, nil
	}

	// gets pairsがあれば，指定分だけ返す
	var s Statuses
	for _, pair := range p.Pairs {
		for _, v := range resp.Data.Statuses {
			if pair == v.Pair {
				s = append(s, v)
			}
		}
	}

	return s, nil
}
