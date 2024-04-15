package orders

import (
	"net/http"

	"github.com/go-numb/go-bitbank/v2/types"
)

type Request struct {
	Pair    string `json:"pair"`
	OrderID int    `json:"order_id"`
}

type Response struct {
	Success int8 `json:"success"`
	Data    Data `json:"data"`
}

type Data struct {
	Code   uint32 `json:"code"`
	Orders Order  `json:"order"`
}

type Order struct {
	OrderID         int64  `json:"order_id"`
	Pair            string `json:"pair"`
	Side            string `json:"side"`
	Type            string `json:"type"`
	StartAmount     string `json:"start_amount"`
	RemainingAmount string `json:"remaining_amount"`
	ExecutedAmount  string `json:"executed_amount"`
	Price           string `json:"price,omitempty"`
	PostOnly        bool   `json:"post_only,omitempty"`
	AveragePrice    string `json:"average_price"`
	OrderedAt       int64  `json:"ordered_at"`
	ExpireAt        int64  `json:"expire_at,omitempty"`
	TriggeredAt     int64  `json:"triggered_at,omitempty"`
	TriggerPrice    string `json:"trigger_price,omitempty"`
	Status          string `json:"status"`
}

func (req *Request) Endpoint() string {
	return types.ENDPOINTPRIVATE
}

func (req *Request) IsAuth() bool {
	return true
}

func (req *Request) Path() string {
	return "user/spot/order"
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
