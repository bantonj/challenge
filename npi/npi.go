package npi

import (
	"fmt"
	"strconv"
)

func getDigits(i int) []int {
	var is []int

	s := strconv.Itoa(i)
	for d := 0; d < len(s); d++ {
		n, err := strconv.Atoi(string(s[d]))
		if err != nil {
			panic(fmt.Sprintf("this shouldn't be possible: %s", err))
		}

		is = append(is, n)
	}

	return is
}

// with assistance from Wikipedia: https://en.wikipedia.org/wiki/Luhn_algorithm
func check(ds []int) bool {
	n := len(ds)
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

// CheckNPIs takes a slice of int and returns a map keyed those ints declaring
// whether or not they are valid NPIs.
func CheckNPIs(npis []int) (map[int]bool, error) {
	results := make(map[int]bool, len(npis))

	for _, npi := range npis {
		results[npi] = check(getDigits(npi))
	}

	return results, nil
}
