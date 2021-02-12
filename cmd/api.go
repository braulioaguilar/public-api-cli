package main

import (
	"fmt"
	"strings"

	"github.com/brauliodev29/public-apis/pkg/presenter"
)

const (
	URL_PREFIX = "https://api.publicapis.org"
)

var (
	data  presenter.Entry
	list  []string
	query string
)

// Build filter string
func getQuery(filter *string) (qs string) {
	filterPart := strings.Split(*filter, ",")
	for _, p := range filterPart {
		qs += fmt.Sprintf("%v&", p)
	}

	return qs[:len(qs)-1]
}

// Run func
func Run() (interface{}, error) {
	if categoriesSubComand.Parsed() {
		if err := httpClient.GetData(fmt.Sprintf("%v/categories", URL_PREFIX), &list); err != nil {
			return nil, err
		}

		return list, nil
	}

	if entriesSubComand.Parsed() {
		if len(*entryFilter) > 0 {
			query = getQuery(entryFilter)
		}

		if err := httpClient.GetData(fmt.Sprintf("%v/entries?%v", URL_PREFIX, query), &data); err != nil {
			return nil, err
		}
	}

	if randomSubComand.Parsed() {
		if len(*randomFilter) > 0 {
			query = getQuery(randomFilter)
		}

		if err := httpClient.GetData(fmt.Sprintf("%v/random?%v", URL_PREFIX, query), &data); err != nil {
			return nil, err
		}
	}

	return data, nil
}
