package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/credsimple/challenge/records"
)

var (
	file = flag.String("file", "", "location of the file containing records")
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

	// TODO(Erik): sorting
	for _, r := range rs {
		fmt.Println(r)
	}
}
