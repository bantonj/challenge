package main

import (
	"flag"
	"fmt"
	"os"
	"sort"
	"strings"

	"github.com/credsimple/challenge/records"
)

type sortableRecords struct {
	rs   []records.Record
	less func([]records.Record, int, int) bool
}

func (r sortableRecords) Len() int {
	return len(r.rs)
}

func (r sortableRecords) Less(i, j int) bool {
	return r.less(r.rs, i, j)
}

func (r sortableRecords) Swap(i, j int) {
	r.rs[i], r.rs[j] = r.rs[j], r.rs[i]
}

func newProviderSorter(rs []records.Record) sortableRecords {
	return sortableRecords{
		rs: rs,
		less: func(rs []records.Record, i, j int) bool {
			if rs[i].ProviderType == rs[j].ProviderType {
				// ascending
				return rs[i].LastName < rs[j].LastName
			}

			// ascending
			return rs[i].ProviderType < rs[j].ProviderType
		},
	}
}

func newDOBSorter(rs []records.Record) sortableRecords {
	return sortableRecords{
		rs: rs,
		less: func(rs []records.Record, i, j int) bool {
			// ascending
			return rs[i].DateOfBirth.Before(rs[j].DateOfBirth)
		},
	}
}

func newLastNameSorter(rs []records.Record) sortableRecords {
	return sortableRecords{
		rs: rs,
		less: func(rs []records.Record, i, j int) bool {
			// descending
			return rs[i].LastName > rs[j].LastName
		},
	}
}

func getSorter(s string, rs []records.Record) (sort.Interface, error) {
	switch s {
	case "provider":
		return newProviderSorter(rs), nil

	case "dob":
		return newDOBSorter(rs), nil

	case "lastname":
		return newLastNameSorter(rs), nil

	default:
		return nil, fmt.Errorf("invalid sort parameter: %s", s)
	}
}

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
