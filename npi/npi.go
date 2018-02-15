package npi

import (
	"fmt"
	"strconv"
)

func getDigits(i int) []int {
	var is []int

	s := strconv.Itoa(i)
	for d := 0; d < len(s); d++ {
		// when you index a string, you get a byte; restring it
		// error is ignored because it should be impossible to fail
		n, _ := strconv.Atoi(string(s[d]))
		is = append(is, n)
	}

	return is
}

// with assistance from Wikipedia: https://en.wikipedia.org/wiki/Luhn_algorithm
func checkLuhn(ds []int) bool {
	n := len(ds)

	// we want to double every other digit starting from the end
	// if n is odd, skip the first digit
	// if n is even, start with the first digit
	parity := n % 2

	var sum int
	for i := 0; i < n; i++ {
		d := ds[i]
		if i%2 == parity {
			d *= 2
		}

		if d > 9 {
			d -= 9
		}

		sum += d
	}

	return sum%10 == 0
}

func getLuhnCheckDigit(ds []int) int {
	var sum int

	double := true
	for i := len(ds) - 1; i >= 0; i-- {
		d := ds[i]
		if double {
			d *= 2
		}

		if d > 9 {
			d -= 9
		}

		sum += d
		double = !double
	}

	sum *= 9
	sumDigits := getDigits(sum)
	return sumDigits[len(sumDigits)-1]
}

const (
	npiPrefix = 80840
)

// CheckNPIs takes a slice of int and returns a map keyed those ints declaring
// whether or not they are valid NPIs.
func CheckNPIs(npis []int) (map[int]bool, error) {
	results := make(map[int]bool, len(npis))

	for _, npi := range npis {
		ds := getDigits(npi)
		l := len(ds)
		ds = append(getDigits(npiPrefix), ds...)

		switch l {
		case 10:
			results[npi] = checkLuhn(ds)

		case 9:
			ds = append(ds, getLuhnCheckDigit(ds))
			results[npi] = checkLuhn(ds)

		default:
			return nil, fmt.Errorf("NPIs should contain 9 or 10 digits: %d (%d)", npi, l)
		}
	}

	return results, nil
}
