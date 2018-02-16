package records

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"time"
)

// Record is a simple data object corresponding to a row.
type Record struct {
	FirstName     string
	MiddleInitial string
	LastName      string
	Gender        string
	ProviderType  string
	DateOfBirth   time.Time
}

const (
	pipeRune  = '|'
	commaRune = ','
	spaceRune = ' '
)

func findDelimiter(bs []byte) (rune, error) {
	for _, r := range []rune{pipeRune, commaRune, spaceRune} {
		if bytes.ContainsRune(bs, r) {
			return r, nil
		}
	}

	return 0x0, errors.New("could not identify delimiter")
}

// ParseRecords take an io.Reader containing records in various formats and
// returns them in a consistent POGO.
func ParseRecords(r io.Reader) ([]Record, error) {
	// there may be a more memory efficient way to do this, but it's less
	// reader-efficient
	bs, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, fmt.Errorf("error reading from reader: %s", err)
	}

	delim, err := findDelimiter(bs)
	if err != nil {
		return nil, err
	}

	parser, err := getParser(delim)
	if err != nil {
		return nil, err
	}

	buf := bytes.NewBuffer(bs)
	return parser.Parse(buf)
}
