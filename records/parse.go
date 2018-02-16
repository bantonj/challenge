package records

import (
	"encoding/csv"
	"fmt"
	"io"
	"strings"
	"time"
)

type fieldParser func(*Record, string) error

func passthrough(r *string, s string) error {
	*r = s
	return nil
}

func firstNameParser(r *Record, s string) error {
	return passthrough(&r.FirstName, s)
}

func lastNameParser(r *Record, s string) error {
	return passthrough(&r.LastName, s)
}

func middleInitialParser(r *Record, s string) error {
	return passthrough(&r.MiddleInitial, s)
}

func providerParser(r *Record, s string) error {
	return passthrough(&r.ProviderType, s)
}

func genderParser(r *Record, s string) error {
	switch strings.ToLower(s) {
	case "m", "male":
		r.Gender = "Male"

	case "f", "female":
		r.Gender = "Female"

	default:
		r.Gender = "Unknown"
	}

	return nil
}

func dobParser(layout string) fieldParser {
	return func(r *Record, s string) error {
		dob, err := time.Parse(layout, s)
		if err != nil {
			return fmt.Errorf("error parsing time (%s): %s", s, err)
		}

		// With two digit years, sometimes Go parses them as 20xx instead of 19xx.
		// 70 *seems* to be the cutoff.
		if dob.After(time.Now()) {
			dob = dob.AddDate(-100, 0, 0)
		}

		r.DateOfBirth = dob
		return nil
	}
}

type recordParser struct {
	delim        rune
	fieldParsers []fieldParser
}

func (p recordParser) Parse(r io.Reader) ([]Record, error) {
	var rs []Record

	rr := csv.NewReader(r)
	rr.Comma = p.delim
	rr.FieldsPerRecord = len(p.fieldParsers)

	for {
		fields, err := rr.Read()
		if err == io.EOF {
			return rs, nil
		}
		if err != nil {
			return nil, fmt.Errorf("error reading input: %s", err)
		}

		var r Record
		for i := 0; i < rr.FieldsPerRecord; i++ {
			p.fieldParsers[i](&r, fields[i])
		}

		rs = append(rs, r)
	}
}

var fieldParsersMap = map[rune][]fieldParser{
	// LastName|FirstName|MiddleInitial|Gender(M/F)|ProviderType|DateOfBirth(M-D-YY)
	pipeRune: []fieldParser{
		lastNameParser,
		firstNameParser,
		middleInitialParser,
		genderParser,
		providerParser,
		dobParser("1-2-06"),
	},

	// LastName,FirstName,Gender(Male/Female),ProviderType,DateOfBirth(M/D/YYYY)
	commaRune: []fieldParser{
		lastNameParser,
		firstNameParser,
		genderParser,
		providerParser,
		dobParser("1/2/2006"),
	},

	// LastName FirstName MiddleInitial Gender(M/F) DateOfBirth(M-D-YYYY) ProviderType
	spaceRune: []fieldParser{
		lastNameParser,
		firstNameParser,
		middleInitialParser,
		genderParser,
		dobParser("1-2-2006"),
		providerParser,
	},
}

func getParser(delim rune) (recordParser, error) {
	p := recordParser{
		delim: delim,
	}

	var ok bool
	if p.fieldParsers, ok = fieldParsersMap[delim]; !ok {
		return p, fmt.Errorf("cannot find parser for delimiter %q", delim)
	}

	return p, nil
}
