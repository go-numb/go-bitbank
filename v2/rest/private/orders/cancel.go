package orders

import (
	"encoding/json"
	"net/http"

	"github.com/go-numb/go-bitbank/v2/types"
)

type RequestForCancel struct {
	Pair    string `json:"pair"`
	OrderID int64  `json:"order_id"`
}

type ResponseForCancel struct {
	Success int8 `json:"success"`
	Data    Data `json:"data"`
}

func (req *RequestForCancel) Endpoint() string {
	return types.ENDPOINTPRIVATE
}

func (req *RequestForCancel) IsAuth() bool {
	return true
}

func (req *RequestForCancel) Path() string {
	return "user/spot/cancel_order"
}

func (req *RequestForCancel) Method() string {
	return http.MethodPost
}

func (req *RequestForCancel) Query() string {
	return ""
}

func (req *RequestForCancel) Payload() []byte {
	b, err := json.Marshal(req)
	if err != nil {
		return nil
	}
	return b
}
