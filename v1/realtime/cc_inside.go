package realtime

import (
	"fmt"

	depth "github.com/go-numb/go-bitbank/v1/depths"
	"github.com/go-numb/go-bitbank/v1/ohlcv"
	transaction "github.com/go-numb/go-bitbank/v1/transactions"

	ticker "github.com/go-numb/go-bitbank/v1/tickers"
	"github.com/gorilla/websocket"
)

type Client struct {
	conn *websocket.Conn

	// old↓
	// Subscriber chan interface{}
	// New↓
	Subscriber chan Recive
	Done       chan error
}

type Types int
type Pairs int

const (
	// Channel prefix
	TypeError Types = iota
	TypeTicker
	TypeDepthAll
	TypeDepthDiff
	TypeTransactions
	TypeCandlestick
)

const (
	// Pairs btc_jpy, xrp_jpy, ltc_btc, eth_btc, mona_jpy, mona_btc, bcc_jpy, bcc_btc
	PairBTCJPY Pairs = iota
	PairXRPJPY
	PairBCCJPY
	PairMNAJPY

	PairETHBTC
	PairLTCBTC
	PairMNABTC
	PairBCCBTC
)

type Recive struct {
	Types Types // int
	Pairs Pairs // int

	Transactions transaction.Transactions
	Depth        depth.Depth
	Tickers      ticker.Ticker
	OHLCV        ohlcv.OHLCVs
	Error        error
}

func (ws *Client) Ping() error {
	pingRegulation := `2`

	err := ws.conn.WriteMessage(websocket.TextMessage, []byte(pingRegulation))
	if err != nil {
		return err
	}

	return nil
}

func (ws *Client) SubscribeWholeDepth(pair string) error {
	b := fmt.Sprintf(`42["join-room", "depth_whole_%s"]`, pair)

	err := ws.conn.WriteMessage(websocket.TextMessage, []byte(b))
	if err != nil {
		return err
	}

	return nil
}

func (ws *Client) SubscribeDiffDepth(pair string) error {
	b := fmt.Sprintf(`42["join-room", "depth_diff_%s"]`, pair)

	err := ws.conn.WriteMessage(websocket.TextMessage, []byte(b))
	if err != nil {
		return err
	}

	return nil
}

func (ws *Client) SubscribeTransactions(pair string) error {
	b := fmt.Sprintf(`42["join-room", "transactions_%s"]`, pair)

	err := ws.conn.WriteMessage(websocket.TextMessage, []byte(b))
	if err != nil {
		return err
	}

	return nil
}

func (ws *Client) SubscribeTicker(pair string) error {
	b := fmt.Sprintf(`42["join-room", "ticker_%s"]`, pair)

	err := ws.conn.WriteMessage(websocket.TextMessage, []byte(b))
	if err != nil {
		return err
	}

	return nil
}
