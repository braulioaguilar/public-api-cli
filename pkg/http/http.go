package http

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"
)

// Request struct
type Request struct {
	Client *http.Client
}

// NewClient func
func NewClient() *Request {
	return &Request{
		Client: &http.Client{
			Timeout: 3 * time.Second,
		},
	}
}

// GetData func
func (r Request) GetData(url string, data interface{}) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	// fill data
	if err := json.Unmarshal(body, &data); err != nil {
		return err
	}

	return nil
}
