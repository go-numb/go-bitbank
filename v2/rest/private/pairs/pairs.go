package pairs

import (
	"net/http"
	"path"

	"github.com/go-numb/go-bitbank/v2/types"
)

type Request struct {
}

type Response struct {
	Success int  `json:"success"`
	Data    Data `json:"data"`
}

type Data struct {
	Pairs []Pair `json:"pairs"`
}

type Pair struct {
	Name                string  `json:"name"`
	BaseAsset           string  `json:"base_asset"`
	QuoteAsset          string  `json:"quote_asset"`
	MakerFeeRateBase    string  `json:"maker_fee_rate_base"`
	TakerFeeRateBase    string  `json:"taker_fee_rate_base"`
	MakerFeeRateQuote   string  `json:"maker_fee_rate_quote"`
	TakerFeeRateQuote   string  `json:"taker_fee_rate_quote"`
	UnitAmount          string  `json:"unit_amount"`
	LimitMaxAmount      string  `json:"limit_max_amount"`
	MarketMaxAmount     string  `json:"market_max_amount"`
	MarketAllowanceRate string  `json:"market_allowance_rate"`
	PriceDigits         float64 `json:"price_digits"`
	AmountDigits        float64 `json:"amount_digits"`
	IsEnabled           bool    `json:"is_enabled"`
	StopOrder           bool    `json:"stop_order"`
	StopOrderAndCancel  bool    `json:"stop_order_and_cancel"`
	StopMarketOrder     bool    `json:"stop_market_order"`
	StopStopOrder       bool    `json:"stop_stop_order"`
	StopStopLimitOrder  bool    `json:"stop_stop_limit_order"`
	StopBuyOrder        bool    `json:"stop_buy_order"`
	StopSellOrder       bool    `json:"stop_sell_order"`
}

func (req *Request) Endpoint() string {
	return types.ENDPOINTPRIVATE
}

func (req *Request) IsAuth() bool {
	return false
}

func (req *Request) Path() string {
	return path.Join("spot", "pairs")
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
