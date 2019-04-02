package realtime

import (
	"bytes"
	"fmt"
	"time"

	"github.com/go-numb/go-bitbank/tickers"

	"github.com/go-numb/go-bitbank/transactions"

	"github.com/go-numb/go-bitbank/depths"

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

func TestRealtime(t *testing.T) {
	c, err := Connect(false)
	if err != nil {
		t.Error(err)
	}

	channels := []string{Depth, Ticker, Transactions}
	pairs := []string{BTCJPY}
	c.SetSubscribes(channels, pairs)
	go c.Realtime(channels, pairs)

	for {
		select {
		case v := <-c.Subscriber:
			switch v.(type) {
			case depth.Depth:
				fmt.Printf("%+v\n", v)
			case depth.DepthDiff:
				fmt.Printf("%+v\n", v)
			case transaction.Transaction:
				fmt.Printf("%+v\n", v)
			case ticker.Ticker:
				fmt.Printf("%+v\n", v)
			}
		}
	}
}
