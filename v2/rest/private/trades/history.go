package trades

import (
	"net/http"
	"time"

	"github.com/go-numb/go-bitbank/v2/types"
	"github.com/google/go-querystring/query"
)

type Request struct {
	Pair      string    `url:"pair"`
	Count     int64     `url:"count,omitempty"`
	OrderID   int64     `url:"order_id,omitempty"`
	Since     time.Time `url:"-"`
	End       time.Time `url:"-"`
	SinceUnix int64     `url:"since,omitempty"`
	EndUnix   int64     `url:"end,omitempty"`
	Order     string    `url:"order,omitempty"`
}

type Response struct {
	Success int8 `json:"success"`
	Data    Data `json:"data"`
}

type Data struct {
	Code   uint32  `json:"code"`
	Trades []Trade `json:"trades"`
}

type Trade struct {
	TradeID        int64      `json:"trade_id"`
	OrderID        int64      `json:"order_id"`
	Pair           string     `json:"pair"`
	Type           string     `json:"type"`
	Side           string     `json:"side"`
	Price          string     `json:"price"`
	Amount         string     `json:"amount"`
	MakerTaker     string     `json:"maker_taker"`
	FeeAmountBase  string     `json:"fee_amount_base"`
	FeeAmountQuote string     `json:"fee_amount_quote"`
	ExecutedAt     types.Time `json:"executed_at"`
}

func (req *Request) Endpoint() string {
	return types.ENDPOINTPRIVATE
}

func (req *Request) IsAuth() bool {
	return true
}

func (req *Request) Path() string {
	return "user/spot/trade_history"
}

func (req *Request) Method() string {
	return http.MethodGet
}

func (req *Request) Query() string {
	t := time.Time{}
	if req.Since != t {
		req.SinceUnix = req.Since.Unix()
	}
	if req.End != t {
		req.EndUnix = req.End.Unix()
	}

	v, err := query.Values(req)
	if err != nil {
		return ""
	}

	return v.Encode()
}

func (req *Request) Payload() []byte {
	return nil
}
