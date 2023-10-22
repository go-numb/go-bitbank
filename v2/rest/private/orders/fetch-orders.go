package orders

import (
	"encoding/json"
	"net/http"

	"github.com/go-numb/go-bitbank/v2/types"
)

type RequestForFetchOrders struct {
	Pair     string `json:"pair"`
	OrderIDs []int  `json:"order_ids"`
}

type ResponseForFetchOrders struct {
	Success int8 `json:"success"`
	Data    Data `json:"data"`
}

type DataForFetchOrders struct {
	Code   uint32                `json:"code"`
	Orders []OrderForCreateOrder `json:"order"`
}

func (req *RequestForFetchOrders) Endpoint() string {
	return types.ENDPOINTPRIVATE
}

func (req *RequestForFetchOrders) Path() string {
	return "user/spot/orders_info"
}

func (req *RequestForFetchOrders) Method() string {
	return http.MethodPost
}

func (req *RequestForFetchOrders) Query() string {
	return ""
}

func (req *RequestForFetchOrders) Payload() []byte {
	b, err := json.Marshal(req)
	if err != nil {
		return nil
	}
	return b
}
