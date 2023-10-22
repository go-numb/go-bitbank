package types

import (
	"encoding/json"
	"strconv"
	"time"
)

type Ohlcv struct {
	Open      float64   `json:"open,string"`
	High      float64   `json:"high,string"`
	Low       float64   `json:"low,string"`
	Close     float64   `json:"close,string"`
	Volume    float64   `json:"volume,string"`
	Timestamp time.Time `json:"timestamp"`
}

func (p *Ohlcv) UnmarshalJSON(b []byte) error {
	var values []any
	if err := json.Unmarshal(b, &values); err != nil {
		return err
	}

	var (
		open, high, low, close, volume float64
		t                              int64
		timestamp                      time.Time
	)

	for i, v := range values {
		switch i {
		case 0:
			open, _ = strconv.ParseFloat(v.(string), 64)
		case 1:
			high, _ = strconv.ParseFloat(v.(string), 64)
		case 2:
			low, _ = strconv.ParseFloat(v.(string), 64)
		case 3:
			close, _ = strconv.ParseFloat(v.(string), 64)
		case 4:
			volume, _ = strconv.ParseFloat(v.(string), 64)
		case 5:
			switch value := v.(type) {
			case string:
				t, _ = strconv.ParseInt(value, 10, 64)
				timestamp = time.Unix(int64(t/1000), 0)
			case int64:
				timestamp = time.Unix(0, int64(time.Duration(value)*time.Millisecond))
			case int:
				timestamp = time.Unix(0, int64(time.Duration(int64(value))*time.Millisecond))
			case float64:
				d := duration(value)
				timestamp = time.Unix(0, int64(d))
			}
		}
	}

	p.Open = open
	p.High = high
	p.Low = low
	p.Close = close
	p.Volume = volume
	p.Timestamp = timestamp

	return nil
}

func duration(f float64) time.Duration {
	return time.Duration(f) * time.Millisecond
}
