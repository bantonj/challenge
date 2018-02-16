package records

import (
	"encoding/csv"
	"fmt"
	"io"
)

type csvParser struct{}

func (p *csvParser) Parse(r io.Reader) ([]Record, error) {
	var rs []Record

	csvR := csv.NewReader(r)
	csvR.FieldsPerRecord = 5

	for {
		fields, err := csvR.Read()
		if err == io.EOF {
			return rs, nil
		}
		if err != nil {
			return nil, fmt.Errorf("error reading input: %s", err)
		}

		// LastName,FirstName,Gender(Male/Female),ProviderType,DateOfBirth(M/D/YYYY)
		r := Record{
			LastName:     fields[0],
			FirstName:    fields[1],
			Gender:       parseGender(fields[2]),
			ProviderType: fields[3],
		}

		dob, err := parseDOB("1/2/2006", fields[4])
		if err != nil {
			return nil, err
		}
		r.DateOfBirth = dob

		rs = append(rs, r)
	}
}
