package records

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"strings"
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

// TODO(Erik): instead of hand-writing individual parsers, what about a config
// style thing?
type parser interface {
	Parse(io.Reader) ([]Record, error)
}

type delimiter int8

const (
	unknown delimiter = iota
	pipe
	comma
	space
)

func getParser(delim delimiter) (parser, error) {
	switch delim {
	case pipe:
		return &psvParser{}, nil

	case comma:
		return &csvParser{}, nil

	case space:
		return &ssvParser{}, nil

	default:
		return nil, errors.New("invalid delimiter")
	}
}

func findDelimiter(bs []byte) (delimiter, error) {
	if bytes.Contains(bs, []byte{'|'}) {
		return pipe, nil
	}

	if bytes.Contains(bs, []byte{','}) {
		return comma, nil
	}

	if bytes.Contains(bs, []byte{' '}) {
		return space, nil
	}

	return unknown, errors.New("could not identify delimiter")
}

// ParseRecords take an io.Reader containing records in various formats and
// returns them in a consistent POGO.
func ParseRecords(r io.Reader) ([]Record, error) {
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

func parseGender(g string) string {
	switch strings.ToLower(g) {
	case "m", "male":
		return "Male"

	case "f", "female":
		return "Female"

	default:
		return "Unknown"
	}
}

// TODO(Erik): fix birthdates later than this year
func parseDOB(layout string, t string) (time.Time, error) {
	dob, err := time.Parse(layout, t)
	if err != nil {
		return time.Time{}, fmt.Errorf("error parsing time (%s): %s", t, err)
	}
	return dob, nil
}
