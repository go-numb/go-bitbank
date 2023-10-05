package ohlcv

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/url"
	"path"
	"strings"
	"time"

	e "github.com/go-numb/go-bitbank/errors"

	"github.com/go-numb/go-bitbank/types"
)

type Timespan int

const (
	BASEURL = "https://public.bitbank.cc/"
	PATH    = "candlestick"

	//  1min 5min 15min 30min 1hour 4hour 8hour 12hour 1day 1week
	Timespan1m Timespan = iota
	Timespan5m
	Timespan15m
	Timespan30m
	Timespan1h
	Timespan4h
	Timespan8h
	Timespan12h
	Timespan1D
	Timespan1W
)

type Request struct {
	// btc_jpy, xrp_jpy, ltc_btc, eth_btc, mona_jpy, mona_btc, bcc_jpy, bcc_btc
	Pair string
	// Timespan as candle-type
	Timespan string
	// YYYY or YYYYMMDD string
	AtDate string
}

type Response struct {
	Success int `json:"success"`
	Data    struct {
		Code        int        `json:"code"`
		Candlestick []OHLCVs   `json:"candlestick"`
		Timestamp   types.Time `json:"timestamp"`
	} `json:"data"`
}

type OHLCVs struct {
	Type   string       `json:"type"`
	OHLCVs types.OHLCVs `json:"ohlcv"`
}

// Set pair string
func (p *Request) Set(pair string, timespan Timespan, t time.Time) {
	p.Pair = strings.ToLower(pair)

	switch timespan { //  1min 5min 15min 30min 1hour 4hour 8hour 12hour 1day 1week
	case Timespan1m:
		p.Timespan = "1min"
	case Timespan5m:
		p.Timespan = "5min"
	case Timespan15m:
		p.Timespan = "15min"
	case Timespan30m:
		p.Timespan = "30min"
	case Timespan1h:
		p.Timespan = "1hour"
	case Timespan4h:
		p.Timespan = "4hour"
	case Timespan8h:
		p.Timespan = "8hour"
	case Timespan12h:
		p.Timespan = "12hour"
	case Timespan1D:
		p.Timespan = "1day"
	case Timespan1W:
		p.Timespan = "1week"
	default:
		p.Timespan = "1day"
	}

	p.AtDate = t.Format("20060102")
}

func (p *Request) Get() (OHLCVs, error) {
	u, err := url.ParseRequestURI(BASEURL)
	if err != nil {
		return OHLCVs{}, err
	}
	u.Path = path.Join(p.Pair, PATH, p.Timespan, p.AtDate)

	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return OHLCVs{}, err
	}

	c := new(http.Client)
	res, err := c.Do(req)
	if err != nil {
		return OHLCVs{}, err
	}
	if res.StatusCode != 200 {
		return OHLCVs{}, errors.New(res.Status)
	}
	defer res.Body.Close()

	var resp Response
	json.NewDecoder(res.Body).Decode(&resp)
	if resp.Success != 1 {
		return OHLCVs{}, e.Handler(resp.Data.Code, err)
	}

	return resp.Data.Candlestick[0], nil
}
