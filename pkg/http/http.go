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

// Options struct
type Options struct {
	Timeout time.Duration
}

// NewClient func
func NewClient(opts *Options) *request {
	if opts == nil {
		opts = &Options{
			Timeout: 3 * time.Second,
		}
	}

	return &request{
		Client: &http.Client{
			Timeout: opts.Timeout,
		},
	}
}

// GetData func
func (r request) GetData(url string, data interface{}) error {
	resp, err := r.Client.Get(url)
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
