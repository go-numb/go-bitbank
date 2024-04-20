package rest

import (
	"fmt"
	"testing"

	"github.com/go-numb/go-bitbank/v2/auth"
	"github.com/go-numb/go-bitbank/v2/rest/private/assets"
	"github.com/go-numb/go-bitbank/v2/rest/private/deposit"
	"github.com/go-numb/go-bitbank/v2/rest/private/orders"
	"github.com/go-numb/go-bitbank/v2/rest/private/pairs"
	"github.com/go-numb/go-bitbank/v2/rest/private/trades"
	"github.com/stretchr/testify/assert"
)

const (
	KEY    = ""
	SECRET = ""
)

func TestPairs(t *testing.T) {
	c := New(nil)
	res, err := c.Pairs(&pairs.Request{})
	if err != nil {
		t.Fatal(err)
	}
	assert.NoError(t, err)

	for _, v := range res.Data.Pairs {
		fmt.Printf("Success: %s, %s, %f, %f\n", v.Name, v.UnitAmount, v.PriceDigits, v.AmountDigits)
	}

	fmt.Printf("Success: %+v\n", res)
}

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

func TestOrder(t *testing.T) {
	c := New(&auth.Config{
		Key:    KEY,
		Secret: SECRET,
	})
	res, err := c.Order(&orders.Request{
		Pair:    "rndr_jpy",
		OrderID: 0,
	})
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
		Pair:     "rndr_jpy",
		Type:     "limit",
		Side:     "buy",
		Amount:   0.01,
		Price:    1240.00,
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
		Pair:    "rndr_jpy",
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
		Pair:     "rndr_jpy",
		OrderIDs: []int64{35247868021, 1},
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
		Pair: "rndr_jpy",
		// Count:  0,
		// FromID: 0,
		// Since:  time.Now().Add(-24 * 30 * time.Hour),
		// End:    time.Now(),
	})
	if err != nil {
		t.Fatal(err)
	}
	assert.NoError(t, err)

	for _, v := range res.Data.Orders {
		fmt.Printf("Success: %d - %f, %s, %v\n", v.OrderID, v.Price, v.Status, v.OrderedAt)
	}
}

func TestFetchTrades(t *testing.T) {
	c := New(&auth.Config{
		Key:    KEY,
		Secret: SECRET,
	})
	res, err := c.FetchTrades(&trades.Request{
		Pair: "rndr_jpy",
		// Count:  0,
		// FromID: 0,
		// Since:  time.Now().Add(-24 * 30 * time.Hour),
		// End:    time.Now(),
	})
	assert.NoError(t, err)

	for _, v := range res.Data.Trades {
		fmt.Printf("Success: %d - %s, %d, %v\n", v.OrderID, v.Price, v.TradeID, v.ExecutedAt)
	}
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
