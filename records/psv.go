package records

import (
	"bufio"
	"fmt"
	"io"
	"strings"
	"time"
)

type psvParser struct{}

func (p *psvParser) Parse(r io.Reader) ([]Record, error) {
	var rs []Record

	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		t := scanner.Text()

		// LastName|FirstName|MiddleInitial|Gender(M/F)|ProviderType|DateOfBirth(M-D-YY)
		parts := strings.Split(t, "|")
		r := Record{
			LastName:      parts[0],
			FirstName:     parts[1],
			MiddleInitial: parts[2],
			Gender:        parseGender(parts[3]),
			ProviderType:  parts[4],
		}

		dob, err := time.Parse("1-2-06", parts[5])
		if err != nil {
			return nil, fmt.Errorf("error parsing time (%s): %s", parts[5], err)
		}
		r.DateOfBirth = dob

		rs = append(rs, r)
	}
	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("error parsing PSV: %s", err)
	}

	return rs, nil
}
