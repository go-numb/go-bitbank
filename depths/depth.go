package depth

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"path"
	"strings"

	"gitlab.com/k-terashima/go-bitbank/types"
)

const (
	BASEURL = "https://public.bitbank.cc/"
	PATH    = "depth"
)

type Request struct {
	// btc_jpy, xrp_jpy, ltc_btc, eth_btc, mona_jpy, mona_btc, bcc_jpy, bcc_btc
	Pair string
}
type Response struct {
	Success int   `json:"success"`
	Data    Depth `json:"data"`
}

type Depth struct {
	Asks      types.Books `json:"asks,string"`
	Bids      types.Books `json:"bids,string"`
	Timestamp types.Time  `json:"timestamp"`
}

// Set pair string
func (p *Request) Set(pair string) {
	p.Pair = strings.ToLower(pair)
}

func (p *Request) Get() (Depth, error) {
	url := BASEURL + path.Join(p.Pair, PATH)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return Depth{}, err
	}

	c := new(http.Client)
	res, err := c.Do(req)
	if err != nil {
		return Depth{}, err
	}
	defer res.Body.Close()

	var resp Response
	json.NewDecoder(res.Body).Decode(&resp)
	if resp.Success != 1 {
		return Depth{}, errors.New(fmt.Sprintf("response error, not success. error code is %d", resp.Success))
	}

	return resp.Data, nil
}
