# go-bitbank

bitbank API

## Description

go-bitbank is a go client library for [Bitbank.cc API](https://docs.bitbank.cc).

## Installation

```
$ go get -u github.com/go-numb/go-bitbank
```

## PublicAPI
``` go
// Versoin1  since2019
package main

import (
 "fmt"
 "github.com/go-numb/go-bitbank"
)


func main() {
	c := bitbank.New("", "")

	c.Ticker.Set("btc_jpy")
	res, err := c.Ticker.Get()
	if err != nil {
		t.Error(err)
	}

	fmt.Printf("%v\n", res)

	doSomething()
}
```

```go 
// Versoin.2 Since2023.11
package main

func main() {
	c := New(nil)
	res, err := c.Depth(&depth.Request{
		Pair: "btc_jpy",
	})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Success: %+v\n", res)

	// 指定枚数以上の板をリストアップ
	depths := res.Data.Wall(5)

	fmt.Println("Ask wall: ")
	for i := 0; i < len(depths.Asks.Books); i++ {
		fmt.Printf("%f: %f\n", depths.Asks.Books[i].Price, depths.Asks.Books[i].Size)
	}
	for i := 0; i < len(depths.Bids.Books); i++ {
		fmt.Printf("%f: %f\n", depths.Bids.Books[i].Price, depths.Bids.Books[i].Size)
	}

}

```

## PrivateAPI

``` go
// Version.1 Since2019
package main

import (
 "fmt"
 "github.com/go-numb/go-bitbank"
)

func main() {

    // loads API keys
	var s Sign
	toml.DecodeFile("../.keys/bitbank.toml", &s)
	c := New(s.Token, s.Secret)

	res, err := c.Auth.Orders.Post(
		&orders.Body{
			Pair:   "bcc_btc",
			Amount: 0.0001, // to string
			Price:  0.04082002,
			Side:   "sell",
			Type:   "limit",
	})
	if err != nil {
		t.Error(err)
	}

	fmt.Printf("%v\n", res)

	doSomething()
}
```

```go 
// Version.2 Since2023.11
package main

func main() {
	c := New(&auth.Config{
		Key:    "[Account APIKEY]",
		Secret: "[Account APISECRETKEY]",
	})
	res, err := c.Assets(&assets.Request{})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Success: %+v\n", res)
}

```

## websocket/realtime
``` go
func main() {
	c, err := realtime.Connect(false)
	if err != nil {
		log.Err(err)
	}

	channels := []string{realtime.ChDepth, realtime.ChTicker, realtime.ChTransactions}
	// Can be specified as a string
	pairs := []string{
		realtime.BTCJPY,
		realtime.XRPJPY,
		realtime.ETHBTC,
	}
	c.SetSubscribes(channels, pairs)
	go c.Realtime(channels, pairs)
	defer c.Close()

EndF:
	for v := range c.Subscriber {
		switch v.Types {
		case realtime.TypeDepthAll:
			fmt.Printf("depth all: %+v\n", v.Depth)
		case realtime.TypeDepthDiff:
			fmt.Printf("depth diff: %+v\n", v.Depth)
		case realtime.TypeTicker:
			fmt.Printf("ticker: %+v\n", v.Tickers)
		case realtime.TypeCandlestick:
			fmt.Printf("candle: %+v\n", v.OHLCV)
		case realtime.TypeTransactions:
			fmt.Printf("transaction: %+v\n", v.Transactions)

		case realtime.TypeError:
			break EndF
		}
	}

	// 実行プリント
	// 	ticker: {Code:0 Sell:4.029001e+06 Buy:4.029e+06 High:4.055e+06 Low:3.946e+06 Last:4.029001e+06 Vol:305.8786 Timestamp:{Time:2023-09-29 14:44:56.608 +0900 JST}}
	// ticker: {Code:0 Sell:75.972 Buy:75.971 High:76.466 Low:74.1 Last:75.973 Vol:5.5717360327e+06 Timestamp:{Time:2023-09-29 14:44:56.176 +0900 JST}}
	// depth diff: {Code:0 Asks:{Books:[]} Bids:{Books:[{Price:4.02701e+06 Size:0}]} Timestamp:{Time:2023-09-29 14:44:58.003 +0900 JST}}

	log.Fatal().Msg("END")
}
```

## Author

[@_numbP](https://twitter.com/_numbP)

## License

[MIT](https://github.com/go-numb/go-bitbank/blob/master/LICENSE)