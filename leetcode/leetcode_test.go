package leetcode

import (
	"math"
	"testing"
)

func TestConvertTemperature(t *testing.T) {
	cases := []struct {
		celsius  float64
		expected []float64
	}{
		{36.50, []float64{309.65, 97.7}},
		{122.11, []float64{395.26, 251.798}},
	}

	epsilon := 0.00001
	for _, c := range cases {
		got := convertTemperature(c.celsius)
		if len(got) != len(c.expected) {
			t.Errorf("convertTemperature(%f) != %v, got %v\n", c.celsius, c.expected, got)
			continue
		}
		for i := range got {
			err := math.Abs(got[i] - c.expected[i])
			if err > epsilon {
				t.Errorf("convertTemperature(%f) != %v, got %v\n", c.celsius, c.expected[i], got[i])
				break
			}
		}
	}
}

func TestMaxFreqSum(t *testing.T) {
	cases := []struct {
		s        string
		expected int
	}{
		{"successes", 6},
		{"aeiaeia", 3},
	}

	for _, c := range cases {
		got := maxFreqSum(c.s)
		if got != c.expected {
			t.Errorf("maxFreqSum(%q) != %d, got %d\n", c.s, c.expected, got)
		}
	}
}

func TestFilterCharacters(t *testing.T) {
	cases := []struct {
		s        string
		k        int
		expected string
	}{
		{"aadbbcccca", 3, "dbb"},
		{"xyz", 2, "xyz"},
	}

	for _, c := range cases {
		got := filterCharacters(c.s, c.k)
		if got != c.expected {
			t.Errorf("filterCharacters(%q, %d) != %q, got %q\n", c.s, c.k, c.expected, got)
		}
	}
}
