package transactions

import (
	"fmt"
	"net/http"
	"path"
	"time"

	"github.com/go-numb/go-bitbank/v2/types"
)

type Request struct {
	// btc_jpy, xrp_jpy, ltc_btc, eth_btc, mona_jpy, mona_btc, bcc_jpy, bcc_btc
	Pair string
	Term time.Time
}

type Response struct {
	Success int  `json:"success"`
	Data    Data `json:"data"`
}

type Data struct {
	Code         uint32        `json:"code"`
	Transactions []Transaction `json:"transactions"`
}

type Transaction struct {
	TransactionId int64      `json:"transaction_id"`
	Side          string     `json:"side"`
	Price         float64    `json:"price,string"`
	Amount        float64    `json:"amount,string"`
	ExecutedAt    types.Time `json:"executed_at"`
}

func (req *Request) Endpoint() string {
	return types.ENDPOINTPUBLIC
}

func (req *Request) IsAuth() bool {
	return false
}

func (req *Request) Path() string {
	term := fmt.Sprintf(
		"%d%02d%02d",
		req.Term.Year(),
		int8(req.Term.Month()),
		req.Term.Day())

	return path.Join(req.Pair, "transactions", term)
}

func (req *Request) Method() string {
	return http.MethodGet
}

func (req *Request) Query() string {
	return ""
}

func (req *Request) Payload() []byte {
	return nil
}
