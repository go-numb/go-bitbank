package depth

import (
	"encoding/json"
	"math"
	"net/http"
	"net/url"
	"path"
	"strings"
	"sync"

	e "github.com/go-numb/go-bitbank/v1/errors"

	"github.com/go-numb/go-bitbank/v1/types"
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
	Code      int         `json:"code"`
	Asks      types.Books `json:"asks,a,string"`
	Bids      types.Books `json:"bids,b,string"`
	Timestamp types.Time  `json:"timestamp,t"`
}

type DepthDiff struct {
	Code      int         `json:"code"`
	Asks      types.Books `json:"a,string"`
	Bids      types.Books `json:"b,string"`
	Timestamp types.Time  `json:"t"`
}

// Set pair string
func (p *Request) Set(pair string) {
	p.Pair = strings.ToLower(pair)
}

func (p *Request) Get() (Depth, error) {
	u, err := url.ParseRequestURI(BASEURL)
	if err != nil {
		return Depth{}, err
	}
	u.Path = path.Join(p.Pair, PATH)

	req, err := http.NewRequest("GET", u.String(), nil)
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
		return Depth{}, e.Handler(resp.Data.Code, err)
	}
	return resp.Data, nil
}

// Aggregate now price * monitoringRangePcnt 範囲の待機板集計
func (p Depth) Aggregate(monitoringRangePcnt float64) (float64, float64) {
	var wg sync.WaitGroup
	var askvol, bidvol float64

	askL := len(p.Asks.Books)
	bidL := len(p.Bids.Books)
	if askL <= 0 || bidL <= 0 { // 素材があるか確認
		return 0, 0
	}

	// 範囲集計の準備，最終約定ではなく均衡金額
	var mid float64
	mid = float64(p.Asks.Books[0].Price+p.Bids.Books[0].Price) / 2

	monitor := mid * monitoringRangePcnt

	wg.Add(1)
	go func() {
		for _, v := range p.Asks.Books {
			askvol += v.Size

			if mid+monitor < v.Price {
				break
			}
		}

		wg.Done()
	}()

	wg.Add(1)
	go func() {
		for _, v := range p.Bids.Books {
			bidvol += v.Size

			if mid-monitor > v.Price {
				break
			}
		}

		wg.Done()
	}()

	wg.Wait()

	return askvol, bidvol
}

// Eat 約定があったものとして現在の待機板を喰ってみる
func (p Depth) Eat(eatAsk, eatBid float64) (bestask int, bestbid int) {
	var wg sync.WaitGroup

	askL := len(p.Asks.Books)
	bidL := len(p.Bids.Books)
	if askL <= 0 || bidL <= 0 { // 素材があるか確認
		return 0, 0
	}

	wg.Add(1)
	go func() {
		for _, v := range p.Asks.Books {
			eatAsk -= v.Size

			if eatAsk < 0 {
				bestask = int(math.RoundToEven(v.Price - 1.0))
				break
			}
		}

		wg.Done()
	}()

	wg.Add(1)
	go func() {
		for _, v := range p.Bids.Books {
			eatBid -= v.Size

			if eatBid < 0 {
				bestbid = int(math.RoundToEven(v.Price + 1.0))
				break
			}
		}

		wg.Done()
	}()

	wg.Wait()

	return bestask, bestbid
}
