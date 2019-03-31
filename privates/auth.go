package private

import (
	"gitlab.com/k-terashima/go-bitbank/privates/assets"
	"gitlab.com/k-terashima/go-bitbank/privates/orders"
	"gitlab.com/k-terashima/go-bitbank/privates/status"
	"gitlab.com/k-terashima/go-bitbank/privates/trades"
)

type Auth struct {
	Status *status.Request
	Assets *assets.Request
	Orders *orders.Request
	Trades *trades.Request
}

func New(token, secret string) *Auth {
	return &Auth{
		Status: &status.Request{
			Token:  token,
			Secret: secret,
		},
		Assets: &assets.Request{
			Token:  token,
			Secret: secret,
		},
		Orders: &orders.Request{
			Token:  token,
			Secret: secret,
		},
		Trades: &trades.Request{
			Token:  token,
			Secret: secret,
		},
	}
}
