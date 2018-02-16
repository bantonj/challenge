package main

import (
	"flag"
	"fmt"
	"os"
	"sort"
	"strings"

	"github.com/credsimple/challenge/records"
)

func recordToString(r records.Record) string {
	// last name, first name, gender, date of birth, provider type. Display dates in the format MM/DD/YYYY.
	return strings.Join(
		[]string{
			r.LastName,
			r.FirstName,
			r.Gender,
			r.DateOfBirth.Format("01/02/2006"),
			r.ProviderType,
		},
		", ",
	)
}

var (
	file    = flag.String("file", "", "location of the file containing records")
	sortOpt = flag.String("sort", "", "sort the records: provider, dob, or lastname")
)

func main() {
	flag.Parse()
	if file == nil || *file == "" {
		flag.Usage()
		os.Exit(1)
	}

	f, err := os.Open(*file)
	if err != nil {
		fmt.Fprintf(os.Stderr, "could not read not file: %s", err)
		os.Exit(1)
	}

	rs, err := records.ParseRecords(f)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error parsing records: %s", err)
		os.Exit(1)
	}

	if sortOpt != nil && *sortOpt != "" {
		sorter, err := getSorter(*sortOpt, rs)
		if err != nil {
			fmt.Fprintf(os.Stderr, "error identifying sort parameter: %s", err)
		}

		sort.Sort(sorter)
	}

	for _, r := range rs {
		fmt.Println(recordToString(r))
	}
}
