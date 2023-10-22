package rest

import (
	"bytes"
	"fmt"
	"net/http"
	"net/url"
	"path"

	"github.com/go-numb/go-bitbank/v2/auth"
	"github.com/rs/zerolog/log"
)

type Requester interface {
	Endpoint() string
	Path() string
	Method() string
	Query() string
	Payload() []byte
}

type Response struct {
	Success int8 `json:"success"`
	Data    struct {
		Code uint32 `json:"code"`
	} `json:"data"`
}

func (p *Client) request(req Requester, results interface{}) error {
	res, err := p.do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if err := json.NewDecoder(res.Body).Decode(results); err != nil {
		log.Debug().Str("decode", err.Error())
		return err
	}

	return nil
}

func (p *Client) newRequest(r Requester) *http.Request {
	// avoid Pointer's butting
	u, _ := url.ParseRequestURI(r.Endpoint())
	u.Path = path.Join(u.Path, r.Path())
	u.RawQuery = r.Query()

	body := r.Payload()
	req, err := http.NewRequest(r.Method(), u.String(), bytes.NewReader(body))
	if err != nil {
		return nil
	}

	if p.Auth != nil {
		auth.MakeHeader(p.Auth.Key, p.Auth.Secret, body, req)
		if req == nil {
			log.Debug().Str("creates request", "error")
			return nil
		}
	}

	return req
}

func (c *Client) do(r Requester) (*http.Response, error) {
	req := c.newRequest(r)

	res, err := c.HTTPC.Do(req)
	if err != nil {
		return nil, err
	}

	if res.StatusCode != 200 || err != nil {
		res.Body.Close()
		if err != nil {
			return nil, &APIError{
				Status: uint32(res.StatusCode),
				Message: fmt.Sprintf(
					"status code: %d, error: %s, status: %s",
					res.StatusCode,
					err.Error(),
					res.Status),
			}
		}

		return nil, &APIError{
			Status: uint32(res.StatusCode),
			Message: fmt.Sprintf(
				"status code: %d, status: %s",
				res.StatusCode,
				res.Status),
		}
	}

	return res, nil
}
