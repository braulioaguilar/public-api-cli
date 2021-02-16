package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/hokaccha/go-prettyjson"
)

var (
	entriesSubCommand    *flag.FlagSet
	randomSubCommand     *flag.FlagSet
	categoriesSubCommand *flag.FlagSet
	entryFilter          *string
	randomFilter         *string
	endpoints            map[string]string
)

// Filter struct
type Filter struct {
	Entry, Random string
}

func main() {
	// Register main endpoints
	endpoints = map[string]string{
		"categories": "/categories",
		"entries":    "/entries",
		"random":     "/random",
	}

	// Create subcomand
	entriesSubCommand = flag.NewFlagSet("entries", flag.ExitOnError)
	randomSubCommand = flag.NewFlagSet("random", flag.ExitOnError)
	categoriesSubCommand = flag.NewFlagSet("categories", flag.ExitOnError)

	entryFilter = entriesSubCommand.String("filter", "", "String query filters")
	randomFilter = randomSubCommand.String("filter", "", "String query filters")
	flag.Parse()

	if len(os.Args) < 2 {
		fmt.Println("We need 2 parameters to continue request.")
		os.Exit(1)
	}

	// Check first argument as endpoint
	source := os.Args[1]
	if _, ok := endpoints[source]; !ok {
		fmt.Println("Argument not valid")
		os.Exit(1)
	}

	switch source {
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

	data, err := Run(&Filter{*entryFilter, *randomFilter})
	if err != nil {
		fmt.Printf("Error: %s\n", err.Error())
		os.Exit(1)
	}

	e, _ := prettyjson.Marshal(data)
	fmt.Printf("%v\n", string(e))
}
