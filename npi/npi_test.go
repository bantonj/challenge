package npi

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCheckNPIs(t *testing.T) {
	t.Parallel()

	tests := []struct {
		input  int
		output bool
	}{
		{
			79927398710,
			false,
		},
		{
			79927398711,
			false,
		},
		{
			79927398712,
			false,
		},
		{
			79927398713,
			true,
		},
		{
			79927398714,
			false,
		},
		{
			79927398715,
			false,
		},
		{
			79927398716,
			false,
		},
		{
			79927398717,
			false,
		},
		{
			79927398718,
			false,
		},
		{
			79927398719,
			false,
		},
	}

	for _, test := range tests {
		test := test

		t.Run(strconv.Itoa(test.input), func(t *testing.T) {
			t.Parallel()

			m := CheckNPIs([]int{test.input})

			require.Contains(t, m, test.input)
			assert.Equal(t, m[test.input], test.output)
		})
	}
}
