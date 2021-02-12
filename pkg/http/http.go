package http

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"
)

// Request struct
type request struct {
	Client *http.Client
}

// NewClient func
func NewClient() *request {
	return &request{
		Client: &http.Client{
			Timeout: 3 * time.Second,
		},
	}
}

// GetData func
func (r request) GetData(url string, data interface{}) error {
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
