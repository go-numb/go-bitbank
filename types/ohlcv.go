package types

import (
	"bytes"
	"encoding/json"
	"time"
)

type OHLCVs struct {
	OHLCVs []OHLCV
}

type OHLCV struct {
	Open     float64
	High     float64
	Low      float64
	Close    float64
	Size     float64
	Timespan time.Time
}

func (p *OHLCVs) UnmarshalJSON(b []byte) error {
	source := bytes.Replace(b, []byte(`"`), []byte(``), -1)

	var m [][]interface{}
	json.Unmarshal(source, &m)

	for _, v := range m {
		t := v[5].(float64)
		tt := int64(t) / int64(time.Microsecond)
		p.OHLCVs = append(p.OHLCVs, OHLCV{
			Open:     v[0].(float64),
			High:     v[1].(float64),
			Low:      v[2].(float64),
			Close:    v[3].(float64),
			Size:     v[4].(float64),
			Timespan: time.Unix(tt, 0),
		})

	}

	return nil
}

// p.Open = m[0].(float64)
// p.High = m[1].(float64)
// p.Low = m[2].(float64)
// p.Close = m[3].(float64)
// p.Size = m[4].(float64)

// t := m[5].(int64)
// t = t / int64(time.Microsecond)
// p.Timespan = time.Unix(t, 0)
