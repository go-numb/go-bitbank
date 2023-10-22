package rest

import (
	"fmt"
	"testing"

	"github.com/go-numb/go-bitbank/v2/auth"
	"github.com/go-numb/go-bitbank/v2/rest/private/assets"
	"github.com/go-numb/go-bitbank/v2/rest/private/deposit"
	"github.com/go-numb/go-bitbank/v2/rest/private/orders"
	"github.com/go-numb/go-bitbank/v2/rest/private/trades"
	"github.com/stretchr/testify/assert"
)

const (
	KEY    = ""
	SECRET = ""
)

func TestAssets(t *testing.T) {
	c := New(&auth.Config{
		Key:    KEY,
		Secret: SECRET,
	})
	res, err := c.Assets(&assets.Request{})
	if err != nil {
		t.Fatal(err)
	}
	assert.NoError(t, err)

	fmt.Printf("Success: %+v\n", res)
}

func TestOrders(t *testing.T) {
	c := New(&auth.Config{
		Key:    KEY,
		Secret: SECRET,
	})
	res, err := c.Orders(&orders.Request{})
	if err != nil {
		t.Fatal(err)
	}
	assert.NoError(t, err)

	fmt.Printf("Success: %+v\n", res)
}

func TestCreateOrder(t *testing.T) {
	c := New(&auth.Config{
		Key:    KEY,
		Secret: SECRET,
	})
	res, err := c.CreateOrder(&orders.RequestForCreateOrder{
		Pair:     "btc_jpy",
		Type:     "limit",
		Side:     "buy",
		Amount:   0.01,
		Price:    4400000,
		PostOnly: true,
	})
	if err != nil {
		t.Fatal(err)
	}
	assert.NoError(t, err)

	fmt.Printf("Success: %+v\n", res)
}

func TestCancel(t *testing.T) {
	c := New(&auth.Config{
		Key:    KEY,
		Secret: SECRET,
	})
	res, err := c.Cancel(&orders.RequestForCancel{
		Pair:    "btc_jpy",
		OrderID: 0,
	})
	if err != nil {
		t.Fatal(err)
	}
	assert.NoError(t, err)

	fmt.Printf("Success: %+v\n", res)
}

func TestFetchOrders(t *testing.T) {
	c := New(&auth.Config{
		Key:    KEY,
		Secret: SECRET,
	})
	res, err := c.FetchOrders(&orders.RequestForFetchOrders{
		Pair:     "btc_jpy",
		OrderIDs: []int{0, 1},
	})
	if err != nil {
		t.Fatal(err)
	}
	assert.NoError(t, err)

	fmt.Printf("Success: %+v\n", res)
}

func TestFetchActiveOrders(t *testing.T) {
	c := New(&auth.Config{
		Key:    KEY,
		Secret: SECRET,
	})
	res, err := c.FetchActiveOrders(&orders.RequestForFetchActiveOrders{
		Pair: "btc_jpy",
		// Count:  0,
		// FromID: 0,
		// Since:  time.Now().Add(-24 * 30 * time.Hour),
		// End:    time.Now(),
	})
	if err != nil {
		t.Fatal(err)
	}
	assert.NoError(t, err)

	fmt.Printf("Success: %+v\n", res)
}

func TestFetchTrades(t *testing.T) {
	c := New(&auth.Config{
		Key:    KEY,
		Secret: SECRET,
	})
	res, err := c.FetchTrades(&trades.Request{
		Pair: "btc_jpy",
		// Count:  0,
		// FromID: 0,
		// Since:  time.Now().Add(-24 * 30 * time.Hour),
		// End:    time.Now(),
	})
	if err != nil {
		t.Fatal(err)
	}
	assert.NoError(t, err)

	fmt.Printf("Success: %+v\n", res)
}

func TestFetchDeposits(t *testing.T) {
	c := New(&auth.Config{
		Key:    KEY,
		Secret: SECRET,
	})
	res, err := c.FetchDeposits(&deposit.Request{
		Asset: "bcc",
		// Count:  0,
		// Since:  time.Now().Add(-24 * 30 * time.Hour),
		// End:    time.Now(),
	})
	if err != nil {
		t.Fatal(err)
	}
	assert.NoError(t, err)

	fmt.Printf("Success: %+v\n", res)
}
