package depth

import (
	"net/http"
	"path"

	"github.com/go-numb/go-bitbank/v2/types"
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
	Code uint32 `json:"code"`

	Asks types.Books `json:"asks"`
	Bids types.Books `json:"bids"`

	Timestamp  types.Time `json:"timestamp"`
	SequenceId string     `json:"sequenceId"`
}

func (req *Request) Endpoint() string {
	return types.ENDPOINTPUBLIC
}

func (req *Request) Path() string {
	return path.Join(req.Pair, "depth")
}

func (req *Request) Method() string {
	return http.MethodGet
}

func (req *Request) Query() string {
	return ""
}

func (req *Request) Payload() []byte {
	return nil
}

// Wall [ask, bid]
func (p *Depth) Wall(sizeThreshold float64) Depth {
	var (
		price, wall float64

		// Results
		depths = Depth{
			Asks: types.Books{
				Books: make([]types.Book, 0),
			},
			Bids: types.Books{
				Books: make([]types.Book, 0),
			},
		}
	)

	for i := 0; i < len(p.Asks.Books); i++ {
		// 値があり、設定閾値以下ならばスルー
		if sizeThreshold > p.Asks.Books[i].Size && price != 0 {
			continue
		}

		if wall < p.Asks.Books[i].Size {
			// Set
			depths.Asks.Books = append(depths.Asks.Books, types.Book{
				Price: p.Asks.Books[i].Price,
				Size:  p.Asks.Books[i].Size,
			})
			price = p.Bids.Books[i].Price
			wall = p.Bids.Books[i].Size
		}
	}

	// Reset
	price = 0
	wall = 0

	for i := 0; i < len(p.Bids.Books); i++ {
		// 値があり、設定閾値以下ならばスルー
		if sizeThreshold > p.Bids.Books[i].Size && price != 0 {
			continue
		}

		if wall < p.Bids.Books[i].Size {
			// Set
			depths.Bids.Books = append(depths.Bids.Books, types.Book{
				Price: p.Bids.Books[i].Price,
				Size:  p.Bids.Books[i].Size,
			})
			price = p.Bids.Books[i].Price
			wall = p.Bids.Books[i].Size
		}
	}

	return depths
}
