package assets

import (
	"encoding/json"
	"net/http"
	"net/url"
	"path"

	e "github.com/go-numb/go-bitbank/v1/errors"

	"github.com/go-numb/go-bitbank/v1/privates/auth"
)

const (
	BASEURL = "https://api.bitbank.cc/"
	VERSION = "v1"
	PATH    = "user/assets"
)

type Request struct {
	Token  string
	Secret string
}

type Response struct {
	Success int `json:"success"`
	Data    struct {
		Code   int    `json:"code"`
		Assets Assets `json:"assets"`
	} `json:"data"`
}

type Assets []Asset

type Asset struct {
	Asset           string  `json:"asset"`
	AmountPrecision int     `json:"amount_precision"`
	OnhandAmount    float64 `json:"onhand_amount,string"`
	LockedAmount    float64 `json:"locked_amount,string"`
	FreeAmount      float64 `json:"free_amount,string"`
	StopDeposit     bool    `json:"stop_deposit"`
	StopWithdrawal  bool    `json:"stop_withdrawal"`
	WithdrawalFee   struct {
		Threshold float64 `json:"threshold,string"`
		Under     float64 `json:"under,string"`
		Over      float64 `json:"over,string"`
	} `json:"withdrawal_fee"`
}

func (p *Request) Get() (Assets, error) {
	u, err := url.ParseRequestURI(BASEURL)
	if err != nil {
		return Assets{}, err
	}
	u.Path = path.Join(VERSION, PATH)

	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return Assets{}, err
	}

	auth.MakeHeader(p.Token, p.Secret, nil, req)

	c := new(http.Client)
	res, err := c.Do(req)
	if err != nil {
		return Assets{}, err
	}
	defer res.Body.Close()

	var resp Response
	json.NewDecoder(res.Body).Decode(&resp)
	if resp.Success != 1 {
		return Assets{}, e.Handler(resp.Data.Code, err)
	}

	return resp.Data.Assets, nil
}
