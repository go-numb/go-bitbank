package orders

import (
	"encoding/json"
	"net/http"

	"github.com/go-numb/go-bitbank/v2/types"
)

type RequestForCancelOrders struct {
	Pair     string `json:"pair"`
	OrderIDs []int  `json:"order_ids"`
}

type ResponseForCancelOrders struct {
	Success int8                `json:"success"`
	Data    DataForCancelOrders `json:"data"`
}

type DataForCancelOrders struct {
	Code   uint32  `json:"code"`
	Orders []Order `json:"order"`
}

func (req *RequestForCancelOrders) Endpoint() string {
	return types.ENDPOINTPRIVATE
}

func (req *RequestForCancelOrders) Path() string {
	return "user/spot/cancel_orders"
}

func (req *RequestForCancelOrders) Method() string {
	return http.MethodPost
}

func (req *RequestForCancelOrders) Query() string {
	return ""
}

func (req *RequestForCancelOrders) Payload() []byte {
	b, err := json.Marshal(req)
	if err != nil {
		return nil
	}
	return b
}
