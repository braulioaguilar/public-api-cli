package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/brauliodev29/public-apis/pkg/http"
	"github.com/hokaccha/go-prettyjson"
)

var (
	entriesSubCommand    *flag.FlagSet
	randomSubCommand     *flag.FlagSet
	categoriesSubCommand *flag.FlagSet
	entryFilter          *string
	randomFilter         *string
	httpClient           *http.Request
)

func main() {
	// Create subcomand
	entriesSubCommand = flag.NewFlagSet("entries", flag.ExitOnError)
	randomSubCommand = flag.NewFlagSet("random", flag.ExitOnError)
	categoriesSubCommand = flag.NewFlagSet("categories", flag.ExitOnError)

	entryFilter = entriesSubCommand.String("filter", "", "String query filters")
	randomFilter = randomSubCommand.String("filter", "", "String query filters")

	flag.Parse()
	if len(os.Args) < 2 {
		os.Exit(1)
	}

	switch os.Args[1] {
	case "categories":
		categoriesSubCommand.Parse(os.Args[2:])
	case "entries":
		entriesSubCommand.Parse(os.Args[2:])
	case "random":
		randomSubCommand.Parse(os.Args[2:])
	default:
		flag.PrintDefaults()
		os.Exit(1)
	}

	// Config client http
	httpClient = http.NewClient()

	data, err := Run()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}

	e, _ := prettyjson.Marshal(data)
	fmt.Printf("%v\n", string(e))
}
