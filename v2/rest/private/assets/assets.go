package assets

import (
	"net/http"

	"github.com/go-numb/go-bitbank/v2/types"
)

type Request struct {
}

type Response struct {
	Success int8 `json:"success"`
	Data    Data `json:"data"`
}

type Data struct {
	Code   uint32  `json:"code"`
	Assets []Asset `json:"assets"`
}

type Asset struct {
	Asset           string  `json:"asset"`
	AmountPrecision int8    `json:"amount_precision"`
	OnhandAmount    float64 `json:"onhand_amount,string"`
	LockedAmount    float64 `json:"locked_amount,string"`
	FreeAmount      float64 `json:"free_amount,string"`
	StopDeposit     bool    `json:"stop_deposit"`
	StopWithdrawal  bool    `json:"stop_withdrawal"`
	WithdrawalFee   any     `json:"withdrawal_fee"`
}

func (req *Request) Endpoint() string {
	return types.ENDPOINTPRIVATE
}

func (req *Request) IsAuth() bool {
	return true
}

func (req *Request) Path() string {
	return "user/assets"
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
