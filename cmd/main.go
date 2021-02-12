package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/brauliodev29/public-apis/pkg/http"
	"github.com/hokaccha/go-prettyjson"
)

var (
	entriesSubComand    *flag.FlagSet
	randomSubComand     *flag.FlagSet
	categoriesSubComand *flag.FlagSet
	entryFilter         *string
	randomFilter        *string
	httpClient          *http.Request
)

func main() {
	// Create subcomand
	entriesSubComand = flag.NewFlagSet("entries", flag.ExitOnError)
	randomSubComand = flag.NewFlagSet("random", flag.ExitOnError)
	categoriesSubComand = flag.NewFlagSet("categories", flag.ExitOnError)

	entryFilter = entriesSubComand.String("filter", "", "String query filters")
	randomFilter = randomSubComand.String("filter", "", "String query filters")

	flag.Parse()
	if len(os.Args) < 2 {
		os.Exit(1)
	}

	switch os.Args[1] {
	case "categories":
		categoriesSubComand.Parse(os.Args[2:])
	case "entries":
		entriesSubComand.Parse(os.Args[2:])
	case "random":
		randomSubComand.Parse(os.Args[2:])
	default:
		flag.PrintDefaults()
		os.Exit(1)
	}

	// Config cliet http
	httpClient = http.NewClient()

	data, err := Run()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}

	e, _ := prettyjson.Marshal(data)
	fmt.Printf("%v\n", string(e))
}
