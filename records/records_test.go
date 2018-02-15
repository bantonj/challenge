package records

import (
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestParseRecords(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name   string
		input  string
		output []Record
	}{
		{
			"PSV",
			"LastName|FirstName|MiddleInitial|M|ProviderType|1-13-70",
			[]Record{
				{
					FirstName:     "FirstName",
					MiddleInitial: "MiddleInitial",
					LastName:      "LastName",
					Gender:        "Male",
					ProviderType:  "ProviderType",
					DateOfBirth:   time.Date(1970, 1, 13, 0, 0, 0, 0, time.UTC),
				},
			},
		},
		{
			"CSV",
			"LastName,FirstName,Female,ProviderType,1/13/1970",
			[]Record{
				{
					FirstName:    "FirstName",
					LastName:     "LastName",
					Gender:       "Female",
					ProviderType: "ProviderType",
					DateOfBirth:  time.Date(1970, 1, 13, 0, 0, 0, 0, time.UTC),
				},
			},
		},
		{
			"SSV",
			"LastName FirstName MiddleInitial F 1-13-1970 ProviderType",
			[]Record{
				{
					FirstName:     "FirstName",
					MiddleInitial: "MiddleInitial",
					LastName:      "LastName",
					Gender:        "Female",
					ProviderType:  "ProviderType",
					DateOfBirth:   time.Date(1970, 1, 13, 0, 0, 0, 0, time.UTC),
				},
			},
		},
	}

	for _, test := range tests {
		test := test

		t.Run(test.name, func(t *testing.T) {
			rs, err := ParseRecords(strings.NewReader(test.input))
			require.NoError(t, err)

			assert.Equal(t, test.output, rs)
		})
	}
}
