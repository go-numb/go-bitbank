package rest

import (
	"github.com/go-numb/go-bitbank/v2/rest/private/assets"
	"github.com/go-numb/go-bitbank/v2/rest/private/deposit"
	"github.com/go-numb/go-bitbank/v2/rest/private/orders"
	"github.com/go-numb/go-bitbank/v2/rest/private/pairs"
	"github.com/go-numb/go-bitbank/v2/rest/private/trades"
)

func (p *Client) Pairs(req *pairs.Request) (*pairs.Response, error) {
	results := new(pairs.Response)
	if err := p.request(req, results); err != nil {
		return nil, err
	}
	return results, nil
}

func (p *Client) Assets(req *assets.Request) (*assets.Response, error) {
	results := new(assets.Response)
	if err := p.request(req, results); err != nil {
		return nil, err
	}

	// StatusCode == 200 かつ err == nilだが、response body内success表記が成功してない場合
	if results.Success != 1 {
		return nil, &APIError{
			Status:  results.Data.Code,
			Message: "response fail",
		}
	}

	return results, nil
}

func (p *Client) Order(req *orders.Request) (*orders.Response, error) {
	results := new(orders.Response)
	if err := p.request(req, results); err != nil {
		return nil, err
	}

	// StatusCode == 200 かつ err == nilだが、response body内success表記が成功してない場合
	if results.Success != 1 {
		return nil, &APIError{
			Status:  results.Data.Code,
			Message: "response fail",
		}
	}

	return results, nil
}

func (p *Client) CreateOrder(req *orders.RequestForCreateOrder) (*orders.ResponseForCreateOrder, error) {
	results := new(orders.ResponseForCreateOrder)
	if err := p.request(req, results); err != nil {
		return nil, err
	}

	// StatusCode == 200 かつ err == nilだが、response body内success表記が成功してない場合
	if results.Success != 1 {
		return nil, &APIError{
			Status:  results.Data.Code,
			Message: "response fail",
		}
	}

	return results, nil
}

func (p *Client) Cancel(req *orders.RequestForCancel) (*orders.ResponseForCancel, error) {
	results := new(orders.ResponseForCancel)
	if err := p.request(req, results); err != nil {
		return nil, err
	}

	// StatusCode == 200 かつ err == nilだが、response body内success表記が成功してない場合
	if results.Success != 1 {
		return nil, &APIError{
			Status:  results.Data.Code,
			Message: "response fail",
		}
	}

	return results, nil
}

func (p *Client) FetchOrders(req *orders.RequestForFetchOrders) (*orders.ResponseForFetchOrders, error) {
	results := new(orders.ResponseForFetchOrders)
	if err := p.request(req, results); err != nil {
		return nil, err
	}

	// StatusCode == 200 かつ err == nilだが、response body内success表記が成功してない場合
	if results.Success != 1 {
		return nil, &APIError{
			Status:  results.Data.Code,
			Message: "response fail",
		}
	}

	return results, nil
}

func (p *Client) FetchActiveOrders(req *orders.RequestForFetchActiveOrders) (*orders.ResponseForFetchActiveOrders, error) {
	results := new(orders.ResponseForFetchActiveOrders)
	if err := p.request(req, results); err != nil {
		return nil, err
	}

	// StatusCode == 200 かつ err == nilだが、response body内success表記が成功してない場合
	if results.Success != 1 {
		return nil, &APIError{
			// Status:  results.Data.Code,
			Message: "response fail",
		}
	}

	return results, nil
}

func (p *Client) FetchTrades(req *trades.Request) (*trades.Response, error) {
	results := new(trades.Response)
	if err := p.request(req, results); err != nil {
		return nil, err
	}

	// StatusCode == 200 かつ err == nilだが、response body内success表記が成功してない場合
	if results.Success != 1 {
		return nil, &APIError{
			Status:  results.Data.Code,
			Message: "response fail",
		}
	}

	return results, nil
}

func (p *Client) FetchDeposits(req *deposit.Request) (*deposit.Response, error) {
	results := new(deposit.Response)
	if err := p.request(req, results); err != nil {
		return nil, err
	}

	// StatusCode == 200 かつ err == nilだが、response body内success表記が成功してない場合
	if results.Success != 1 {
		return nil, &APIError{
			Status:  results.Data.Code,
			Message: "response fail",
		}
	}

	return results, nil
}
