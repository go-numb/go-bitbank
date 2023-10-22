package types

import (
	"bytes"
	"encoding/json"
)

type Books struct {
	Books []Book
}

type Book struct {
	Price float64
	Size  float64
}

func (p *Books) UnmarshalJSON(b []byte) error {
	source := bytes.Replace(b, []byte(`"`), []byte(``), -1)

	var m [][]float64
	json.Unmarshal(source, &m)

	for _, v := range m {
		p.Books = append(p.Books, Book{ // 4687187371 ns/op
			Price: v[0],
			Size:  v[1],
		})
	}

	return nil
}
