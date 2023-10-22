package rest

import (
	"github.com/go-numb/go-bitbank/v2/rest/public/candle"
	"github.com/go-numb/go-bitbank/v2/rest/public/depth"
	"github.com/go-numb/go-bitbank/v2/rest/public/ticker"
	"github.com/go-numb/go-bitbank/v2/rest/public/transactions"
)

func (p *Client) Depth(req *depth.Request) (*depth.Response, error) {
	results := new(depth.Response)
	if err := p.request(req, results); err != nil {
		return nil, err
	}
	return results, nil
}

func (p *Client) Ticker(req *ticker.Request) (*ticker.Response, error) {
	results := new(ticker.Response)
	if err := p.request(req, results); err != nil {
		return nil, err
	}
	return results, nil
}

func (p *Client) Tickers(req *ticker.RequestForTickers) (*ticker.ResponseForTickers, error) {
	results := new(ticker.ResponseForTickers)
	if err := p.request(req, results); err != nil {
		return nil, err
	}
	return results, nil
}

func (p *Client) TickersJPY(req *ticker.RequestForTickersJPY) (*ticker.ResponseForTickersJPY, error) {
	results := new(ticker.ResponseForTickersJPY)
	if err := p.request(req, results); err != nil {
		return nil, err
	}
	return results, nil
}

func (p *Client) Candle(req *candle.Request) (*candle.Response, error) {
	results := new(candle.Response)
	if err := p.request(req, results); err != nil {
		return nil, err
	}
	return results, nil
}

func (p *Client) Transactions(req *transactions.Request) (*transactions.Response, error) {
	results := new(transactions.Response)
	if err := p.request(req, results); err != nil {
		return nil, err
	}
	return results, nil
}
