package records

import (
	"bufio"
	"fmt"
	"io"
	"strings"
	"time"
)

type csvParser struct{}

func (p *csvParser) Parse(r io.Reader) ([]Record, error) {
	var rs []Record

	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		t := scanner.Text()

		// LastName,FirstName,Gender(Male/Female),ProviderType,DateOfBirth(M/D/YYYY)
		parts := strings.Split(t, ",")
		r := Record{
			LastName:     parts[0],
			FirstName:    parts[1],
			Gender:       parseGender(parts[2]),
			ProviderType: parts[3],
		}

		dob, err := time.Parse("1/2/2006", parts[4])
		if err != nil {
			return nil, fmt.Errorf("error parsing time (%s): %s", parts[4], err)
		}
		r.DateOfBirth = dob

		rs = append(rs, r)
	}
	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("error parsing CSV: %s", err)
	}

	return rs, nil
}
