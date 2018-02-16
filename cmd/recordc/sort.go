package main

import (
	"fmt"
	"sort"

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
