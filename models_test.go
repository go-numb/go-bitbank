package bitbank

import (
	"fmt"
	"strings"
	"testing"
	"time"

	"github.com/go-numb/go-bitbank/privates/orders"

	"github.com/go-numb/go-bitbank/privates/trades"

	"github.com/BurntSushi/toml"

	"github.com/go-numb/go-bitbank/ohlcv"

	"gonum.org/v1/gonum/stat"
)

/* # Public Test
- Transactions: 約定
- Depth: Books/Orderbook
- Ticker: ...
*/
func TestTransactions(t *testing.T) {
	c := New("", "")

	c.Transactions.Set("btc_jpy", time.Now().Add(-10*24*time.Hour))
	res, err := c.Transactions.Get()
	if err != nil {
		t.Error(err)
	}

	// fmt.Printf("%v\n", res)

	var buysideV, sellsideV float64
	var buyAvg, sellAvg []float64
	for _, v := range res {
		if strings.ToLower(v.Side) == "buy" {
			buysideV += v.Amount
			buyAvg = append(buyAvg, v.Price)
		} else {
			sellsideV += v.Amount
			sellAvg = append(sellAvg, v.Price)
		}
	}

	avgB := stat.Mean(buyAvg, nil)
	avgS := stat.Mean(sellAvg, nil)
	fmt.Printf("count: %d, %+v to %+v\n", len(res), res[0].ExecutedAt.Time.String(), res[len(res)-1].ExecutedAt.Time.String())
	fmt.Printf("%.3f - %.1f\n", avgB, buysideV)
	fmt.Printf("%.3f - %.1f\n", avgS, sellsideV)
}

func BenchmarkDepth(t *testing.B) {
	c := New("", "")

	c.Depth.Set("btc_jpy")
	_, err := c.Depth.Get()
	if err != nil {
		t.Error(err)
	}

	// fmt.Printf("%+v\n", res)
}

func TestTicker(t *testing.T) {
	c := New("", "")

	c.Ticker.Set("btc_jpy")
	res, err := c.Ticker.Get()
	if err != nil {
		t.Error(err)
	}

	fmt.Printf("%v\n", res)
}

func TestOHLCV(t *testing.T) {
	c := New("", "")

	c.OHLCV.Set("btc_jpy", ohlcv.Timespan1m, time.Now().Add(-24*time.Hour))
	res, err := c.OHLCV.Get()
	if err != nil {
		t.Error(err)
	}

	fmt.Printf("%v\n", res)
}

type Sign struct {
	Token  string
	Secret string
}

func TestStatus(t *testing.T) {
	var s Sign
	toml.DecodeFile("../.keys/bitbank.toml", &s)
	c := New(s.Token, s.Secret)

	c.Auth.Status.Set()
	res, err := c.Auth.Status.Get()
	if err != nil {
		t.Error(err)
	}

	fmt.Printf("%v\n", res)
}

func TestAssets(t *testing.T) {
	var s Sign
	toml.DecodeFile("../.keys/bitbank.toml", &s)
	c := New(s.Token, s.Secret)

	res, err := c.Auth.Assets.Get()
	if err != nil {
		t.Error(err)
	}

	fmt.Printf("%v\n", res)
}

func TestTrades(t *testing.T) {
	var s Sign
	toml.DecodeFile("../.keys/bitbank.toml", &s)
	c := New(s.Token, s.Secret)

	c.Auth.Trades.Set(&trades.Setter{
		Pair: "btc_jpy",
		// Count: 10.0,
		// Since: time.Now().Add(-48 * time.Hour).Unix(),
		// End:   time.Now().Unix(),
	})
	res, err := c.Auth.Trades.Get()
	if err != nil {
		t.Error(err)
	}

	fmt.Printf("%v\n", res)
}

func TestOrdersPost(t *testing.T) {
	var s Sign
	toml.DecodeFile("../.keys/bitbank.toml", &s)
	c := New(s.Token, s.Secret)

	res, err := c.Auth.Orders.Post(&orders.Body{
		Pair:   "bcc_btc",
		Amount: 0.0001,
		Price:  0.04082002,
		Side:   "sell",
		Type:   "limit",
	})
	if err != nil {
		t.Error(err)
	}

	fmt.Printf("%v\n", res)
}
func TestCancelOrder(t *testing.T) {
	var s Sign
	toml.DecodeFile("../.keys/bitbank.toml", &s)
	c := New(s.Token, s.Secret)

	res, err := c.Auth.Orders.Cancel("bcc_btc", 44326945)
	if err != nil {
		t.Error(err)
	}

	fmt.Printf("%v\n", res)
}

func TestCancelOrders(t *testing.T) {
	var s Sign
	toml.DecodeFile("../.keys/bitbank.toml", &s)
	c := New(s.Token, s.Secret)

	res, err := c.Auth.Orders.Cancels("bcc_btc", 44327083, 44327082)
	if err != nil {
		t.Error(err)
	}

	fmt.Printf("%v\n", res)
}
