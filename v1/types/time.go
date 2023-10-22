package types

import (
	"strconv"
	"time"
)

type Time struct {
	Time time.Time
}

func (p *Time) UnmarshalJSON(b []byte) error {
	s := string(b)
	i, err := strconv.Atoi(s)
	if err != nil {
		return err
	}
	p.Time = time.Unix(int64(i)/1e3, (int64(i)%1e3)*int64(time.Millisecond)/int64(time.Nanosecond))

	return nil
}
