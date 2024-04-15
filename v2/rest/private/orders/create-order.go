package orders

import (
	"encoding/json"
	"net/http"

	"github.com/go-numb/go-bitbank/v2/types"
)

type RequestForCreateOrder struct {
	Pair         string  `json:"pair"`
	Type         string  `json:"type"`
	Side         string  `json:"side"`
	Price        float64 `json:"price"`
	TriggerPrice float64 `json:"trigger_price,omitempty"`
	Amount       float64 `json:"amount"`
	PostOnly     bool    `json:"post_only,omitempty"`
}

type ResponseForCreateOrder struct {
	Success int8               `json:"success"`
	Data    WrapForCreateOrder `json:"data"`
}

type WrapForCreateOrder struct {
	Code   uint32              `json:"code"`
	Orders OrderForCreateOrder `json:"order"`
}

type OrderForCreateOrder struct {
	OrderID int64  `json:"order_id"`
	Pair    string `json:"pair"`
	Side    string `json:"side"`
	Type    string `json:"type"`

	StartAmount     float64 `json:"start_amount,string"`
	RemainingAmount float64 `json:"remaining_amount,string"`
	ExecutedAmount  float64 `json:"executed_amount,string"`
	Price           float64 `json:"price,string"`
	AveragePrice    float64 `json:"average_price,string"`
	TriggerPrice    float64 `json:"trigger_price,string"`
	PostOnly        bool    `json:"post_only"`

	Status      string     `json:"status"`
	OrderedAt   types.Time `json:"ordered_at,omitempty"`
	ExpireAt    types.Time `json:"expire_at,omitempty"`
	CanceledAt  types.Time `json:"canceled_at,omitempty"`
	TriggeredAt types.Time `json:"triggered_at,omitempty"`
}

func (req *RequestForCreateOrder) Endpoint() string {
	return types.ENDPOINTPRIVATE
}

func (req *RequestForCreateOrder) IsAuth() bool {
	return true
}

func (req *RequestForCreateOrder) Path() string {
	return "user/spot/order"
}

func (req *RequestForCreateOrder) Method() string {
	return http.MethodPost
}

func (req *RequestForCreateOrder) Query() string {
	return ""
}

func (req *RequestForCreateOrder) Payload() []byte {
	b, err := json.Marshal(req)
	if err != nil {
		return nil
	}
	return b
}
