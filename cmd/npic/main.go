package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"

	"github.com/credsimple/challenge/npi"
)

var (
	file = flag.String("file", "", "location of the file containing newline-delimited NPIs")
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

	var npis []int
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		t := scanner.Text()
		npi, err := strconv.Atoi(t)
		if err != nil {
			fmt.Fprintf(os.Stderr, "text %q is not valid: %s", t, err)
			os.Exit(1)
		}

		npis = append(npis, npi)
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "error scanning file: %s", err)
		os.Exit(1)
	}

	results, err := npi.CheckNPIs(npis)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error while checking NPIs: %s", err)
		os.Exit(1)
	}

	for npi, valid := range results {
		fmt.Printf("%d: %v\n", npi, valid)
	}
}
