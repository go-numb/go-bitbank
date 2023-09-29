package realtime

import (
	"bytes"
	"fmt"
	"time"

	"github.com/buger/jsonparser"

	"testing"

	"github.com/labstack/gommon/log"

	"github.com/gorilla/websocket"
)

func TestConnect(t *testing.T) {
	var (
	// ticker1m = time.NewTicker(time.Minute)
	)

	isPublic := false
	p, _ := Connect(isPublic)

	pair := "btc_jpy"

	p.SubscribeWholeDepth(pair)
	p.SubscribeDiffDepth(pair)
	p.SubscribeTransactions(pair)
	p.SubscribeTicker(pair)

	go func() {
		ticker20s := time.NewTicker(20 * time.Second)
		for {
			select {
			case <-ticker20s.C:
				if err := p.Ping(); err != nil {
					log.Error(err)
				}
			}
		}
	}()

	paths := [][]string{
		[]string{"room_name"},
		[]string{"message", "data"},
	}

	for {
		_, msg, err := p.conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("error: %v", err)
			}
			break
		}

		switch {
		case bytes.HasPrefix(msg, []byte(`42`)):
			msg = bytes.TrimLeft(msg, `42["message",`)
			msg = bytes.TrimRight(msg, `]`)
			// fmt.Println("42: ", string(msg))
		default:
			fmt.Println("default: ", string(msg))
		}

		// m, vt, n, err := jsonparser.Get(msg, "data")
		// if err != nil {
		// 	log.Error(err)
		// }
		// fmt.Printf("%+v: %d, %s\n", vt, n, string(m))

		jsonparser.EachKey(msg, func(idx int, value []byte, vt jsonparser.ValueType, err error) {
			switch idx {
			case 0: // room_name
				fmt.Printf("%d: type: %v:: %+v\n", idx, vt, string(value))
			case 1: // []string{"person", "avatars", "[0]", "url"}
				// var d depth.Depth
				// json.Unmarshal(value, &d)

				fmt.Printf("%d: type: %v:: %+v\n", idx, vt, string(msg))
			case 2: // []string{"company", "url"},
				fmt.Printf("%d: type: %v:: %+v\n", idx, vt, string(value))
			default:
				fmt.Printf("%d: type: %v:: %+v\n", idx, vt, string(value))
			}
			// fmt.Printf("%s, %v\n", string(msg), err)
		}, paths...)
	}
}

// アーカイブ
func TestRealtime(t *testing.T) {
	c, err := Connect(false)
	if err != nil {
		t.Error(err)
	}

	channels := []string{ChDepth, ChTicker, ChTransactions}
	pairs := []string{
		BTCJPY,
		XRPJPY,
		ETHBTC,
	}
	c.SetSubscribes(channels, pairs)
	go c.Realtime(channels, pairs)
	defer c.Close()

EndF:
	for v := range c.Subscriber {
		switch v.Types {
		case TypeDepthAll:
			fmt.Printf("depth all: %+v\n", v.Depth)
		case TypeDepthDiff:
			fmt.Printf("depth diff: %+v\n", v.Depth)
		case TypeTicker:
			fmt.Printf("ticker: %+v\n", v.Tickers)
		case TypeCandlestick:
			fmt.Printf("candle: %+v\n", v.OHLCV)
		case TypeTransactions:
			fmt.Printf("transaction: %+v\n", v.Transactions)

		case TypeError:
			break EndF
		}
	}

	// 実行プリント
	// 	ticker: {Code:0 Sell:4.029001e+06 Buy:4.029e+06 High:4.055e+06 Low:3.946e+06 Last:4.029001e+06 Vol:305.8786 Timestamp:{Time:2023-09-29 14:44:56.608 +0900 JST}}
	// ticker: {Code:0 Sell:75.972 Buy:75.971 High:76.466 Low:74.1 Last:75.973 Vol:5.5717360327e+06 Timestamp:{Time:2023-09-29 14:44:56.176 +0900 JST}}
	// depth diff: {Code:0 Asks:{Books:[]} Bids:{Books:[{Price:4.02701e+06 Size:0}]} Timestamp:{Time:2023-09-29 14:44:58.003 +0900 JST}}

	log.Fatal("END")
}
