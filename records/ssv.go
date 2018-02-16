package records

import (
	"encoding/csv"
	"fmt"
	"io"
)

type ssvParser struct{}

func (p *ssvParser) Parse(r io.Reader) ([]Record, error) {
	var rs []Record

	ssvR := csv.NewReader(r)
	ssvR.Comma = ' '
	ssvR.FieldsPerRecord = 6

	for {
		fields, err := ssvR.Read()
		if err == io.EOF {
			return rs, nil
		}
		if err != nil {
			return nil, fmt.Errorf("error reading input: %s", err)
		}

		// LastName FirstName MiddleInitial Gender(M/F) DateOfBirth(M-D-YYYY) ProviderType
		r := Record{
			LastName:      fields[0],
			FirstName:     fields[1],
			MiddleInitial: fields[2],
			Gender:        parseGender(fields[3]),
			ProviderType:  fields[5],
		}

		dob, err := parseDOB("1-2-2006", fields[4])
		if err != nil {
			return nil, err
		}
		r.DateOfBirth = dob

		rs = append(rs, r)
	}
}
