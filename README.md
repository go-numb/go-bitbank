# go-bitbank

bitbank API

## Description

go-bitbank is a go client library for [Bitbank.cc API](https://docs.bitbank.cc).

## Installation

```
$ go get -u github.com/go-numb/go-bitbank
```

## PublicAPI
``` 
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

## PrivateAPI

```
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

	res, err := c.Auth.Orders.Post(&orders.Body{
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

## websocket/realtime
```
func main() {
Reconnect:
	wsError := make(chan error)
	
	c, err := realtime.Connect(false)
	if err != nil {
		t.Error(err)
	}

	channels := []string{
		realtime.Depth, 
		realtime.Ticker,
		realtime.Transactions,
	}
	pairs := []string{
		realtime.BTCJPY,
		realtime.XRPJPY,
		realtime.ETHBTC,
	}

	// 購読チャンネルをbitbankに通知し購読開始
	go c.Realtime(channels, pairs)

	for {
		select {
		case v := <-ws.Subscriber:
			switch v.Types {
			case realtime.TypeDepthAll:
				// e.B.Set(true, v.Depth)
				// fmt.Println("depth all\n")
			case realtime.TypeDepthDiff:
				go e.B.Set(false, v.Depth)
				// e.BestPrice()
				// fmt.Printf("diff: %+v\n", v)
				// fmt.Println("depth diff")

			case realtime.TypeTransactions:
				// go e.Set(v.Transactions)
				// fmt.Printf("%.f, delay: %v\n", e.LTP, e.Delay)
				go fmt.Printf("transaction: LTP %f\n", v.Transactions[0].Price)
			case realtime.TypeTicker:
			case realtime.TypeError:
				wsError <- v.Error
			}
		}
	}

	err = <-wsError
	log.Error(err)
	ws.Close()
	time.Sleep(3 * time.Second)
	goto Reconnect
}
```

## Author

[@_numbP](https://twitter.com/_numbP)

## License

[MIT](https://github.com/go-numb/go-bitbank/blob/master/LICENSE)