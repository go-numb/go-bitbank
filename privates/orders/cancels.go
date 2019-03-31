package orders

import (
	"net/url"

	e "github.com/go-numb/go-bitbank/errors"

	"encoding/json"
	"fmt"
	"net/http"
	"path"
	"strings"

	"github.com/go-numb/go-bitbank/privates/auth"
)

func (p *Request) Cancel(pair string, orderID int) (Order, error) {
	u, err := url.ParseRequestURI(BASEURL)
	if err != nil {
		return Order{}, err
	}
	u.Path = path.Join(VERSION, PATH, "cancel_order")

	m := fmt.Sprintf(`{"pair": "%s", "order_id": %d}`, pair, orderID)

	req, err := http.NewRequest("POST", u.String(), strings.NewReader(m))
	if err != nil {
		return Order{}, err
	}

	auth.MakeHeader(p.Token, p.Secret, m, req)

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

func (p *Request) Cancels(pair string, orders ...int) (Order, error) {
	u, err := url.ParseRequestURI(BASEURL)
	if err != nil {
		return Order{}, err
	}
	u.Path = path.Join(VERSION, PATH, "cancel_orders")

	var ids string
	for i, id := range orders {
		if i != 0 {
			ids += ","
		}
		ids += fmt.Sprintf("%d", id)
	}

	m := fmt.Sprintf(`{"pair":"%s","order_ids":[%s]}`, pair, ids)

	req, err := http.NewRequest("POST", u.String(), strings.NewReader(m))
	if err != nil {
		return Order{}, err
	}

	auth.MakeHeader(p.Token, p.Secret, m, req)

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
