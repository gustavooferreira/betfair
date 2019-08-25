package utils

import (
	"fmt"
	"testing"
)

func TestOddExists(t *testing.T) {
	oddsTC := []struct {
		in  int
		out float64
	}{
		{0, 1.01},
		{9, 1.10},
		{99, 2.0},
		{300, 510},
	}

	OddsLen := len(Odds)

	for i, tc := range oddsTC {
		t.Run(fmt.Sprintf("Test %d", i+1), func(t *testing.T) {

			var value float64

			if tc.in >= 0 && tc.in < OddsLen {
				value = Odds[tc.in]
			} else {
				t.Fatalf("Odd index out of range - Number of odds: %d", OddsLen)
			}

			if value != tc.out {
				t.Fatalf("Got '%.2f' odd in position '%d', wanted '%.2f' odd", value, tc.in, tc.out)
			}
		})
	}
}

func TestOddRange(t *testing.T) {

	oddsRandeTC := []struct {
		in       int
		outBegin float64
		outEnd   float64
	}{
		{0, 1.01, 2.0},
		{3, 4.1, 6.0},
		{9, 110, 1000},
	}

	OddsRangeLen := len(OddsRange)

	for i, tc := range oddsRandeTC {
		t.Run(fmt.Sprintf("Test %d", i+1), func(t *testing.T) {

			var value map[string]float64

			if tc.in >= 0 && tc.in < OddsRangeLen {
				value = OddsRange[tc.in]
			} else {
				t.Fatalf("OddRange index out of range - Number of odd ranges: %d", OddsRangeLen)
			}

			begin := value["begin"]
			end := value["end"]

			if begin != tc.outBegin || end != tc.outEnd {
				t.Fatalf("Got begin: '%.2f' and end: '%.2f', wanted begin: '%.2f' and end: '%.2f'", begin, end, tc.outBegin, tc.outEnd)
			}
		})
	}
}

func TestOddFloor(t *testing.T) {
	oddsTC := []struct {
		in  float64
		out float64
	}{
		// Boundary tests
		{-1, 1.01},
		{0, 1.01},
		{1.01, 1.01},
		{99999, 1000},
		{1000, 1000},

		{2.0, 2.0},
		{4.05, 4.00},
		{33, 32},
	}

	for i, tc := range oddsTC {
		t.Run(fmt.Sprintf("Test %d", i+1), func(t *testing.T) {

			value := OddFloor(tc.in)

			if value != tc.out {
				t.Fatalf("Got '%.2f' odd, wanted '%.2f' odd", value, tc.out)
			}
		})
	}
}

func TestOddCeil(t *testing.T) {
	oddsTC := []struct {
		in  float64
		out float64
	}{
		// Boundary tests
		{-1, 1.01},
		{0, 1.01},
		{1.01, 1.01},
		{99999, 1000},
		{1000, 1000},

		{2.0, 2.0},
		{4.05, 4.1},
		{33, 34},
	}

	for i, tc := range oddsTC {
		t.Run(fmt.Sprintf("Test %d", i+1), func(t *testing.T) {

			value := OddCeil(tc.in)

			if value != tc.out {
				t.Fatalf("Got '%.2f' odd, wanted '%.2f' odd", value, tc.out)
			}
		})
	}
}
