package deposit

import (
	"net/http"
	"time"

	"github.com/go-numb/go-bitbank/v2/types"
	"github.com/google/go-querystring/query"
)

type Request struct {
	Asset     string    `url:"asset"`
	Count     int       `url:"count,omitempty"`
	Since     time.Time `url:"-"`
	End       time.Time `url:"-"`
	SinceUnix int64     `url:"since,omitempty"`
	EndUnix   int64     `url:"end,omitempty"`
}

type Response struct {
	Success int8 `json:"success"`
	Data    Data `json:"data"`
}

type Data struct {
	Code     uint32    `json:"code"`
	Deposits []Deposit `json:"deposits"`
}

type Deposit struct {
	UUID        string `json:"uuid"`
	Asset       string `json:"asset"`
	Amount      string `json:"amount"`
	Txid        string `json:"txid"`
	Status      string `json:"status"`
	FoundAt     int    `json:"found_at"`
	ConfirmedAt int    `json:"confirmed_at"`
}

func (req *Request) Endpoint() string {
	return types.ENDPOINTPRIVATE
}

func (req *Request) IsAuth() bool {
	return true
}

func (req *Request) Path() string {
	return "user/deposit_history"
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
