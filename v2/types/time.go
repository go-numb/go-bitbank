package types

import (
	"strconv"
	"time"
)

type Time struct {
	Time time.Time
}

func (p *Time) UnmarshalJSON(b []byte) error {
	i, err := strconv.ParseInt(string(b), 10, 64)
	if err != nil {
		return err
	}
	p.Time = time.Unix(0, int64(time.Duration(i)*time.Millisecond))

	return nil
}
