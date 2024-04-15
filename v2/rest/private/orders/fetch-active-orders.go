package orders

import (
	"net/http"
	"time"

	"github.com/go-numb/go-bitbank/v2/types"
	"github.com/google/go-querystring/query"
)

type RequestForFetchActiveOrders struct {
	Pair      string    `url:"pair"`
	Count     int64     `url:"count,omitempty"`
	FromID    int64     `url:"from_id,omitempty"`
	Since     time.Time `url:"-"`
	End       time.Time `url:"-"`
	SinceUnix int64     `url:"since,omitempty"`
	EndUnix   int64     `url:"end,omitempty"`
}

type ResponseForFetchActiveOrders struct {
	Success int8                     `json:"success"`
	Data    DataForFetchActiveOrders `json:"data"`
}

type DataForFetchActiveOrders struct {
	Code   uint32                `json:"code"`
	Orders []OrderForCreateOrder `json:"order"`
}

func (req *RequestForFetchActiveOrders) Endpoint() string {
	return types.ENDPOINTPRIVATE
}

func (req *RequestForFetchActiveOrders) IsAuth() bool {
	return true
}

func (req *RequestForFetchActiveOrders) Path() string {
	return "user/spot/active_orders"
}

func (req *RequestForFetchActiveOrders) Method() string {
	return http.MethodGet
}

func (req *RequestForFetchActiveOrders) Query() string {
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

func (req *RequestForFetchActiveOrders) Payload() []byte {
	return nil
}
