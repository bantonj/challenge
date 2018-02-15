package records

import (
	"bufio"
	"fmt"
	"io"
	"strings"
	"time"
)

type ssvParser struct{}

// TODO(Erik): bad input breaks this easily
func (p *ssvParser) Parse(r io.Reader) ([]Record, error) {
	var rs []Record

	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		t := scanner.Text()

		// LastName FirstName MiddleInitial Gender(M/F) DateOfBirth(M-D-YYYY) ProviderType
		parts := strings.Split(t, " ")
		r := Record{
			LastName:      parts[0],
			FirstName:     parts[1],
			MiddleInitial: parts[2],
			Gender:        parseGender(parts[3]),
			ProviderType:  parts[5],
		}

		dob, err := time.Parse("1-2-2006", parts[4])
		if err != nil {
			return nil, fmt.Errorf("error parsing time (%s): %s", parts[4], err)
		}
		r.DateOfBirth = dob

		rs = append(rs, r)
	}
	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("error parsing SSV: %s", err)
	}

	return rs, nil
}
