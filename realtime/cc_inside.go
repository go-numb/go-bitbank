package realtime

import (
	"fmt"

	"github.com/gorilla/websocket"
)

type Client struct {
	conn *websocket.Conn

	Subscriber chan interface{}
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
