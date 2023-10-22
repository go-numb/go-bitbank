package ticker

import (
	"net/http"

	"github.com/go-numb/go-bitbank/v2/types"
)

type RequestForTickers struct {
}

type ResponseForTickers struct {
	Success int     `json:"success"`
	Data    Tickers `json:"data"`
}

type Tickers []TickerWithSymbol

type TickerWithSymbol struct {
	Pair string `json:"pair"`
	Ticker
}

func (req *RequestForTickers) Endpoint() string {
	return types.ENDPOINTPUBLIC
}

func (req *RequestForTickers) Path() string {
	return "tickers"
}

func (req *RequestForTickers) Method() string {
	return http.MethodGet
}

func (req *RequestForTickers) Query() string {
	return ""
}

func (req *RequestForTickers) Payload() []byte {
	return nil
}
