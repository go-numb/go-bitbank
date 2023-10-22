package ticker

import (
	"net/http"

	"github.com/go-numb/go-bitbank/v2/types"
)

type RequestForTickersJPY struct {
}

type ResponseForTickersJPY struct {
	Success int     `json:"success"`
	Data    Tickers `json:"data"`
}

func (req *RequestForTickersJPY) Endpoint() string {
	return types.ENDPOINTPUBLIC
}

func (req *RequestForTickersJPY) Path() string {
	return "tickers_jpy"
}

func (req *RequestForTickersJPY) Method() string {
	return http.MethodGet
}

func (req *RequestForTickersJPY) Query() string {
	return ""
}

func (req *RequestForTickersJPY) Payload() []byte {
	return nil
}
