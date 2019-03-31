package transaction

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"path"
	"strings"
	"time"

	"gitlab.com/k-terashima/go-bitbank/types"
)

const (
	BASEURL = "https://public.bitbank.cc/"
	PATH    = "transactions"
)

type Request struct {
	// btc_jpy, xrp_jpy, ltc_btc, eth_btc, mona_jpy, mona_btc, bcc_jpy, bcc_btc
	Pair string
	// YYYMMDD string
	AtDate string
}
type Response struct {
	Success int `json:"success"`
	Data    struct {
		Code         int
		Transactions Transactions `json:"transactions"`
	} `json:"data"`
}

type Transactions []Transaction

type Transaction struct {
	TransactionID int        `json:"transaction_id"`
	Side          string     `json:"side"`
	Price         float64    `json:"price,string"`
	Amount        float64    `json:"amount,string"`
	ExecutedAt    types.Time `json:"executed_at"`
}

// Set pair string
func (p *Request) Set(pair string, op ...interface{}) {
	p.Pair = strings.ToLower(pair)

	for _, o := range op {
		switch v := o.(type) {
		case string:
			p.AtDate = v
		case time.Time:
			p.AtDate = v.Format("20060102")

		}
	}

}

func (p *Request) Get() (Transactions, error) {
	url := BASEURL + path.Join(p.Pair, PATH, p.AtDate)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	c := new(http.Client)
	res, err := c.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	var resp Response
	json.NewDecoder(res.Body).Decode(&resp)
	if resp.Success != 1 {
		return nil, errors.New(fmt.Sprintf("response error, not success. error code is %d", resp.Data.Code))
	}

	return resp.Data.Transactions, nil
}
