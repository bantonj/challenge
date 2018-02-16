package records

import (
	"encoding/csv"
	"fmt"
	"io"
)

type psvParser struct{}

func (p *psvParser) Parse(r io.Reader) ([]Record, error) {
	var rs []Record

	psvR := csv.NewReader(r)
	psvR.Comma = '|'
	psvR.FieldsPerRecord = 6

	for {
		fields, err := psvR.Read()
		if err == io.EOF {
			return rs, nil
		}
		if err != nil {
			return nil, fmt.Errorf("error reading input: %s", err)
		}

		// LastName|FirstName|MiddleInitial|Gender(M/F)|ProviderType|DateOfBirth(M-D-YY)
		r := Record{
			LastName:      fields[0],
			FirstName:     fields[1],
			MiddleInitial: fields[2],
			Gender:        parseGender(fields[3]),
			ProviderType:  fields[4],
		}

		dob, err := parseDOB("1-2-06", fields[5])
		if err != nil {
			return nil, err
		}
		r.DateOfBirth = dob

		rs = append(rs, r)
	}
}
