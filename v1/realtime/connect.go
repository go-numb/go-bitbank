package realtime

import (
	"bytes"
	"encoding/json"

	"github.com/pkg/errors"

	"fmt"
	"time"

	"github.com/labstack/gommon/log"

	"net/http"
	"net/url"
	"strings"

	depth "github.com/go-numb/go-bitbank/v1/depths"
	ticker "github.com/go-numb/go-bitbank/v1/tickers"
	transaction "github.com/go-numb/go-bitbank/v1/transactions"

	"github.com/buger/jsonparser"
	"github.com/gorilla/websocket"
)

const (
	HeartbeatIntervalSecond time.Duration = 25
	ReadTimeoutSecond       time.Duration = 300
	WriteTimeoutSecond      time.Duration = 5
)

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
	head.Set("User-Agent", "Go client")
	head.Set("Cache-Control", "no-cache")
	head.Set("Pragma", "no-cache")
	// head.Set("Sec-WebSocket-Version", "13")

	conn, _, err := d.Dial(u.String(), head)
	if err != nil {
		return nil, err
	}

	return &Client{
		conn: conn,

		// old
		// Subscriber: make(chan interface{}),
		// new
		Subscriber: make(chan Recive),
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
	ChTicker       = "ticker_"
	ChDepth        = "depth_"
	ChTransactions = "transactions_"
	ChCandlestick  = "candlestick_"

	// Channel depth all or diff
	DepthAll  = "whole_"
	DepthDiff = "diff_"

	// Pairs btc_jpy, xrp_jpy, ltc_btc, eth_btc, mona_jpy, mona_btc, bcc_jpy, bcc_btc
	BTCJPY  = "btc_jpy"
	XRPJPY  = "xrp_jpy"
	LTCJPY  = "ltc_jpy"
	BCCJPY  = "bcc_jpy"
	ADAJPY  = "ada_jpy"
	LINKJPY = "link_jpy"
	MATCJPY = "matic_jpy"
	MNAJPY  = "mona_jpy"

	ETHBTC  = "eth_btc"
	XRPBTC  = "xrp_btc"
	LTCBTC  = "ltc_btc"
	BCCBTC  = "bcc_btc"
	LINKBTC = "link_btc"
	MATCBTC = "matic_btc"
	MNABTC  = "mona_btc"
)

// SetSubscribes sets to ws write
func (ws *Client) SetSubscribes(chs, pairs []string) error {
	for _, v := range chs {
		switch {
		case strings.HasPrefix(ChTicker, v):
			for _, pair := range pairs {
				if err := ws.SubscribeTicker(pair); err != nil {
					log.Fatal(err)
				}
				time.Sleep(10 * time.Millisecond)
			}
		case strings.HasPrefix(ChDepth, v):
			for _, pair := range pairs {
				if err := ws.SubscribeWholeDepth(pair); err != nil {
					log.Fatal(err)
				}
				time.Sleep(10 * time.Millisecond)
				if err := ws.SubscribeDiffDepth(pair); err != nil {
					log.Fatal(err)
				}
				time.Sleep(10 * time.Millisecond)
			}
		case strings.HasPrefix(ChTransactions, v):
			for _, pair := range pairs {
				if err := ws.SubscribeTransactions(pair); err != nil {
					log.Fatal(err)
				}
				time.Sleep(10 * time.Millisecond)
			}

		default:
			return errors.New("Unexpected, can not set channnels")
		}
	}

	return nil
}

func (p *Client) Realtime(chs, pairs []string) {
	done := make(chan error)
	fmt.Printf("subscribes: %+v,%+v\n", chs, pairs)
	if err := p.SetSubscribes(chs, pairs); err != nil {
		p.Subscriber <- Recive{
			Types: TypeError,
			Error: err,
		}
	}

	go func() {
		var tickerPong = time.NewTicker(HeartbeatIntervalSecond * time.Second)
		defer tickerPong.Stop()
		for {
			<-tickerPong.C
			p.conn.SetWriteDeadline(time.Now().Add(WriteTimeoutSecond * time.Second))
			if err := p.Ping(); err != nil {
				done <- err
			}
		}
	}()

	go func() {
		for {
			p.conn.SetReadDeadline(time.Now().Add(ReadTimeoutSecond * time.Second))
			_, msg, err := p.conn.ReadMessage()
			if err != nil {
				done <- err
			}

			channelName, _, _, err := jsonparser.Get(msg, "[1]", "room_name")
			if err != nil {
				continue
			}

			switch {
			case bytes.HasPrefix(channelName, []byte(ChDepth+DepthAll)):
				v, _, _, err := jsonparser.Get(msg, "[1]", "message", "data")
				if err != nil {
					continue
				}
				// fmt.Printf("DepthAll: %s\n", string(v))
				var res depth.Depth
				json.Unmarshal(v, &res)
				// fmt.Printf("DepthAll: %+v\n", res)
				p.Subscriber <- Recive{
					Types: TypeDepthAll,
					Pairs: channelNameToPairs(channelName),
					Depth: res,
				}

			case bytes.HasPrefix(channelName, []byte(ChDepth+DepthDiff)):
				v, _, _, err := jsonparser.Get(msg, "[1]", "message", "data")
				if err != nil {
					continue
				}
				// fmt.Printf("DepthDiff: %s\n", string(v))
				var res depth.DepthDiff
				json.Unmarshal(v, &res)
				// fmt.Printf("DepthDiff: %+v\n", res)
				p.Subscriber <- Recive{
					Types: TypeDepthDiff,
					Pairs: channelNameToPairs(channelName),
					Depth: depth.Depth{
						Asks:      res.Asks,
						Bids:      res.Bids,
						Timestamp: res.Timestamp,
					},
				}

			case bytes.HasPrefix(channelName, []byte(ChTransactions)):
				v, _, _, err := jsonparser.Get(msg, "[1]", "message", "data", "transactions")
				if err != nil {
					continue
				}
				// fmt.Printf("Transactions: %s\n", string(v))
				var res transaction.Transactions
				json.Unmarshal(v, &res)
				// fmt.Printf("Transactions: %+v\n", res)
				p.Subscriber <- Recive{
					Types:        TypeTransactions,
					Pairs:        channelNameToPairs(channelName),
					Transactions: res,
				}

			case bytes.HasPrefix(channelName, []byte(ChTicker)):
				v, _, _, err := jsonparser.Get(msg, "[1]", "message", "data")
				if err != nil {
					continue
				}
				// fmt.Printf("Ticker: %s\n", string(v))
				var res ticker.Ticker
				json.Unmarshal(v, &res)
				// fmt.Printf("Ticker: %+v\n", res)
				p.Subscriber <- Recive{
					Types:   TypeTicker,
					Pairs:   channelNameToPairs(channelName),
					Tickers: res,
				}

			default: // 基本的には見ない
				log.Debug("undefined data ------------------------------------------------------------------------------------------------------------")

				p.Subscriber <- Recive{
					Types: TypeError,
					Error: errors.New("undifined error"),
				}
			}
		}
	}()

	err := <-done
	// conform channnel close
	p.Subscriber <- Recive{
		Types: TypeError,
		Error: errors.Wrap(err, "websocket error, "),
	}
}

func channelNameToPairs(name []byte) Pairs {
	switch {
	case bytes.HasSuffix(name, []byte(BTCJPY)):
		return PairBTCJPY
	case bytes.HasSuffix(name, []byte(XRPJPY)):
		return PairXRPJPY
	case bytes.HasSuffix(name, []byte(BCCJPY)):
		return PairBCCJPY
	case bytes.HasSuffix(name, []byte(MNAJPY)):
		return PairMNAJPY
	case bytes.HasSuffix(name, []byte(ETHBTC)):
		return PairETHBTC
	case bytes.HasSuffix(name, []byte(LTCBTC)):
		return PairLTCBTC
	case bytes.HasSuffix(name, []byte(MNABTC)):
		return PairMNABTC
	case bytes.HasSuffix(name, []byte(BCCBTC)):
		return PairBCCBTC
	}

	return PairBTCJPY
}
