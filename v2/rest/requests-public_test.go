package rest

import (
	"fmt"
	"testing"
	"time"

	"github.com/dustin/go-humanize"
	"github.com/go-numb/go-bitbank/v2/rest/public/candle"
	"github.com/go-numb/go-bitbank/v2/rest/public/depth"
	"github.com/go-numb/go-bitbank/v2/rest/public/ticker"
	"github.com/go-numb/go-bitbank/v2/rest/public/transactions"
	"github.com/go-numb/go-bitbank/v2/types"
	"github.com/stretchr/testify/assert"
)

func TestDepth(t *testing.T) {
	c := New(nil)
	res, err := c.Depth(&depth.Request{
		Pair: "btc_jpy",
	})
	if err != nil {
		t.Fatal(err)
	}
	assert.NoError(t, err)

	fmt.Printf("Success: %+v\n", res)

	depths := res.Data.Wall(5)

	fmt.Println("Ask wall: ")
	for i := 0; i < len(depths.Asks.Books); i++ {
		fmt.Printf("%f: %f\n", depths.Asks.Books[i].Price, depths.Asks.Books[i].Size)
	}

	fmt.Println("次Bid wall: ")
	for i := 0; i < len(depths.Bids.Books); i++ {
		fmt.Printf("%f: %f\n", depths.Bids.Books[i].Price, depths.Bids.Books[i].Size)
	}

	fmt.Println("\nask wall: ")

	for i := 0; i < len(depths.Asks.Books); i++ {
		fmt.Printf("%f: %f\n", depths.Asks.Books[i].Price, depths.Asks.Books[i].Size)
	}
	fmt.Println("bids wall: ")
	for i := 0; i < len(depths.Bids.Books); i++ {
		fmt.Printf("%f: %f\n", depths.Bids.Books[i].Price, depths.Bids.Books[i].Size)
	}
}

func TestTicker(t *testing.T) {
	c := New(nil)
	res, err := c.Ticker(&ticker.Request{
		Pair: "btc_jpy",
	})
	if err != nil {
		t.Fatal(err)
	}
	assert.NoError(t, err)

	fmt.Printf("Success: %+v\n", res)
}

func TestTickers(t *testing.T) {
	c := New(nil)
	res, err := c.Tickers(&ticker.RequestForTickers{})
	if err != nil {
		t.Fatal(err)
	}
	assert.NoError(t, err)

	for i := 0; i < len(res.Data); i++ {
		if 0 >= res.Data[i].Vol {
			continue
		}
		fmt.Printf("%d-%s: %+v\n", i, res.Data[i].Pair, res.Data[i])
		fmt.Printf("%s\n", humanize.Commaf(res.Data[i].Last*res.Data[i].Vol))
	}

	fmt.Println("")
	fmt.Println("Volume: 0")
	fmt.Println("")

	for i := 0; i < len(res.Data); i++ {
		if 0 < res.Data[i].Vol {
			continue
		}
		fmt.Printf("%d-%s: %+v\n", i, res.Data[i].Pair, res.Data[i])
		fmt.Printf("%s\n", humanize.Commaf(res.Data[i].Last*res.Data[i].Vol))
	}
}

func TestTickersJPY(t *testing.T) {
	c := New(nil)
	res, err := c.TickersJPY(&ticker.RequestForTickersJPY{})
	if err != nil {
		t.Fatal(err)
	}
	assert.NoError(t, err)

	fmt.Println(len(res.Data))

	for i := 0; i < len(res.Data); i++ {
		if 0 >= res.Data[i].Vol {
			continue
		}
		fmt.Printf("%d-%s: %+v\n", i, res.Data[i].Pair, res.Data[i])
		fmt.Printf("%s円換算\n", humanize.Commaf(res.Data[i].Last*res.Data[i].Vol))
	}

	fmt.Println("")
	fmt.Println("Volume: 0")
	fmt.Println("")
	// JPYでリクエストするとVOlが0のPairが取得できない

	for i := 0; i < len(res.Data); i++ {
		if 0 < res.Data[i].Vol {
			continue
		}
		fmt.Printf("%d-%s: %+v\n", i, res.Data[i].Pair, res.Data[i])
		fmt.Printf("%s円換算\n", humanize.Commaf(res.Data[i].Last*res.Data[i].Vol))
	}
}

func TestCandle(t *testing.T) {
	c := New(nil)
	res, err := c.Candle(&candle.Request{
		Pair:       "btc_jpy",
		CandleType: types.CandleTypeHour1,
		Term:       time.Now(),
	})
	if err != nil {
		t.Fatal(err)
	}
	assert.NoError(t, err)

	fmt.Printf("Success: %+v\n", res)

	for i, v := range res.Data.Candlestick[0].Ohlcvs {
		fmt.Printf("%d - %v\n", i, v)
	}
}

func TestTransactions(t *testing.T) {
	c := New(nil)
	res, err := c.Transactions(&transactions.Request{
		Pair: "btc_jpy",
		Term: time.Now().Add(-time.Hour * 24 * 31),
	})
	if err != nil {
		t.Fatal(err)
	}
	assert.NoError(t, err)

	fmt.Printf("Success: %+v\n", res)
}
