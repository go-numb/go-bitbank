package types

import (
	"encoding/json"
	"strconv"
)

type Books struct {
	Books []Book
}

type Book struct {
	Price float64
	Size  float64
}

func (p *Books) UnmarshalJSON(b []byte) error {
	var values [][]string
	if err := json.Unmarshal(b, &values); err != nil {
		return err
	}

	var (
		price, size float64
	)

	for _, v := range values {
		price, _ = strconv.ParseFloat(v[0], 64)
		size, _ = strconv.ParseFloat(v[1], 64)
		p.Books = append(p.Books, Book{
			Price: price,
			Size:  size,
		})
	}

	return nil
}
