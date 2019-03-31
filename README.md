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

## Author

[@_numbP](https://twitter.com/_numbP)

## License

[MIT](https://github.com/go-numb/go-bitbank/blob/master/LICENSE)