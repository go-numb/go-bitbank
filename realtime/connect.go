package realtime

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/labstack/gommon/log"

	"net/http"
	"net/url"
	"strings"

	depth "github.com/go-numb/go-bitbank/depths"
	ticker "github.com/go-numb/go-bitbank/tickers"
	"github.com/go-numb/go-bitbank/transactions"

	"github.com/buger/jsonparser"
	"github.com/gorilla/websocket"
)

type Request struct {
	subscribeChannels []string
}

func Connect(isPublic bool) (*Client, error) {
	if isPublic {
		return nil, errors.New("not set")
	}

	u := url.URL{
		Scheme:   "wss",
		Host:     "stream.bitbank.cc",
		Path:     "/socket.io/",
		RawQuery: "EIO=3&transport=websocket",
	}

	d := websocket.Dialer{
		Subprotocols:    []string{},
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}
	head := http.Header{"Accept-Encoding": []string{"gzip"}}

	conn, _, err := d.Dial(u.String(), head)
	if err != nil {
		return nil, err
	}

	return &Client{
		conn:       conn,
		Subscriber: make(chan interface{}),
	}, nil
}

func (p *Client) Close() error {
	if err := p.conn.Close(); err != nil {
		return err
	}

	return nil
}

const (
	// Channel prefix
	Ticker       = "ticker_"
	Depth        = "depth_"
	Transactions = "transactions_"
	Candlestick  = "candlestick_"

	// Channel depth all or diff
	DepthAll  = "whole_"
	DepthDiff = "diff_"

	// Pairs btc_jpy, xrp_jpy, ltc_btc, eth_btc, mona_jpy, mona_btc, bcc_jpy, bcc_btc
	BTCJPY = "btc_jpy"
	XRPJPY = "xrp_jpy"
	BCCJPY = "bcc_jpy"
	MNAJPY = "mona_jpy"

	ETHBTC = "eth_btc"
	LTCBTC = "ltc_btc"
	MNABTC = "mona_btc"
	BCCBTC = "bcc_btc"
)

// SetSubscribes sets to ws write
func (ws *Client) SetSubscribes(chs, pairs []string) error {
	for _, v := range chs {
		switch {
		case strings.HasPrefix(Ticker, v):
			for _, pair := range pairs {
				ws.SubscribeTicker(pair)
			}
		case strings.HasPrefix(Depth, v):
			for _, pair := range pairs {
				ws.SubscribeWholeDepth(pair)
				ws.SubscribeDiffDepth(pair)
			}
		case strings.HasPrefix(Transactions, v):
			for _, pair := range pairs {
				ws.SubscribeTransactions(pair)
			}

		default:
			return errors.New("Unexpected, can not set channnels")
		}
	}

	return nil
}

func (p *Client) Realtime(channels, pairs []string) error {
	done := make(chan error)

	go func() {
		for {
			_, msg, err := p.conn.ReadMessage()
			if err != nil {
				if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
					log.Printf("error: %v", err)
				}
				break
			}

			channelName, _, _, err := jsonparser.Get(msg, "[1]", "room_name")
			if err != nil {
				continue
			}

			switch {
			case bytes.HasPrefix(channelName, []byte(Depth+DepthAll)):
				v, _, _, err := jsonparser.Get(msg, "[1]", "message", "data")
				if err != nil {
					continue
				}
				// fmt.Printf("DepthAll: %s\n", string(v))
				var res depth.Depth
				json.Unmarshal(v, &res)
				// fmt.Printf("DepthAll: %+v\n", res)
				p.Subscriber <- res

			case bytes.HasPrefix(channelName, []byte(Depth+DepthDiff)):
				v, _, _, err := jsonparser.Get(msg, "[1]", "message", "data")
				if err != nil {
					continue
				}
				// fmt.Printf("DepthDiff: %s\n", string(v))
				var res depth.DepthDiff
				json.Unmarshal(v, &res)
				// fmt.Printf("DepthDiff: %+v\n", res)
				p.Subscriber <- res

			case bytes.HasPrefix(channelName, []byte(Transactions)):
				v, _, _, err := jsonparser.Get(msg, "[1]", "message", "data", "transactions")
				if err != nil {
					continue
				}
				// fmt.Printf("Transactions: %v(%d): %s\n", vt, n, string(v))
				var res transaction.Transactions
				json.Unmarshal(v, &res)
				fmt.Printf("Transactions: %+v\n", res)
				p.Subscriber <- res

			case bytes.HasPrefix(channelName, []byte(Ticker)):
				v, _, _, err := jsonparser.Get(msg, "[1]", "message", "data")
				if err != nil {
					continue
				}
				// fmt.Printf("Ticker: %s\n", string(v))
				var res ticker.Ticker
				json.Unmarshal(v, &res)
				// fmt.Printf("Ticker: %+v\n", res)
				p.Subscriber <- res

			default: // 基本的には見ない
				log.Debug("------------------------------------------------------\n------------------------------------------------------")
				done <- errors.New("undefined data")
			}
		}
	}()

	return <-done
}
