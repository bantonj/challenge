package npi

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_checkLuhn(t *testing.T) {
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

			valid := checkLuhn(getDigits(test.input))
			assert.Equal(t, valid, test.output)
		})
	}
}

func Test_getLuhnCheckDigit(t *testing.T) {
	t.Parallel()

	tests := []struct {
		input  int
		output int
	}{
		{
			7992739871,
			3,
		},
		{
			123456789,
			7,
		},
		{
			80840123456789,
			3,
		},
	}

	for _, test := range tests {
		test := test

		t.Run(strconv.Itoa(test.input), func(t *testing.T) {
			t.Parallel()

			d := getLuhnCheckDigit(getDigits(test.input))
			assert.Equal(t, d, test.output)
		})
	}
}

func TestCheckNPIs(t *testing.T) {
	t.Parallel()

	tests := []struct {
		input  int
		output bool
	}{
		{
			123456789,
			true,
		},
		{
			1234567893,
			true,
		},
	}

	for _, test := range tests {
		test := test

		t.Run(strconv.Itoa(test.input), func(t *testing.T) {
			t.Parallel()

			m, err := CheckNPIs([]int{test.input})
			require.NoError(t, err)

			require.Contains(t, m, test.input)
			assert.Equal(t, test.output, m[test.input])
		})
	}

	t.Run("long", func(t *testing.T) {
		npi := 12345678901
		_, err := CheckNPIs([]int{npi})
		require.Error(t, err)
	})

	t.Run("short", func(t *testing.T) {
		npi := 12345678
		_, err := CheckNPIs([]int{npi})
		require.Error(t, err)
	})
}
