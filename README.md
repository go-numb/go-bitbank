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
	c, err := realtime.Connect(false)
	if err != nil {
		t.Error(err)
	}

	channels := []string{
		realtime.Depth, 
		ealtime.Ticker,
		realtime.Transactions,
	}
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
			case transaction.Transactions:
				fmt.Printf("%+v\n", v)
			case ticker.Ticker:
				fmt.Printf("%+v\n", v)

			case error:
				goto EndF
			}
		}
	}

EndF:
	c.Close()
}
```

## Author

[@_numbP](https://twitter.com/_numbP)

## License

[MIT](https://github.com/go-numb/go-bitbank/blob/master/LICENSE)