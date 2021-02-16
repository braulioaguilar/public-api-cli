package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/brauliodev29/public-apis/pkg/http"
	"github.com/brauliodev29/public-apis/pkg/presenter"
)

const (
	UrlPrefiX = "https://api.publicapis.org"
)

// Build filter string
func getQuery(filter string) (qs string) {
	filterPart := strings.Split(filter, ",")
	for _, p := range filterPart {
		qs += fmt.Sprintf("%s&", p)
	}

	return qs[:len(qs)-1]
}

// Run func
func Run(f *Filter) (interface{}, error) {
	var (
		data  presenter.Entry
		query string
	)

	// Config client http
	httpClient := http.NewClient(&http.Options{
		Timeout: 5 * time.Second,
	})

	if categoriesSubCommand.Parsed() {
		var list []string
		if err := httpClient.GetData(fmt.Sprintf("%s/categories", UrlPrefiX), &list); err != nil {
			return nil, err
		}

		return list, nil
	}

	if entriesSubCommand.Parsed() {
		if len(f.Entry) > 0 {
			query = getQuery(f.Entry)
		}

		if err := httpClient.GetData(fmt.Sprintf("%s/entries?%s", UrlPrefiX, query), &data); err != nil {
			return nil, err
		}
	}

	if randomSubCommand.Parsed() {
		if len(f.Random) > 0 {
			query = getQuery(f.Random)
		}

		if err := httpClient.GetData(fmt.Sprintf("%s/random?%s", UrlPrefiX, query), &data); err != nil {
			return nil, err
		}
	}

	return data, nil
}
