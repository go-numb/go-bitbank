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
	t := int64(i) / int64(time.Microsecond)
	p.Time = time.Unix(t, 0)

	return nil
}
