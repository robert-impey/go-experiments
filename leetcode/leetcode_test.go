package leetcode

import (
	"math"
	"strconv"
	"strings"
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

func TestAlternatingSum(t *testing.T) {
	cases := []struct {
		nums     []int
		expected int
	}{
		{[]int{1, 3, 5, 7}, -4},
		{[]int{100}, 100},
	}

	for _, c := range cases {
		got := alternatingSum(c.nums)
		if got != c.expected {
			// Convert ints to strings
			strNums := make([]string, len(c.nums))
			for i, v := range c.nums {
				strNums[i] = strconv.Itoa(v)
			}

			// Join the strings with a comma
			numsStr := strings.Join(strNums, ", ")

			t.Errorf("alternatingSum(%s) != %d, got %d\n", numsStr, c.expected, got)
		}
	}
}

func TestFindClosest(t *testing.T) {
	cases := []struct {
		x, y, z  int
		expected int
	}{
		{2, 7, 4, 1},
		{2, 5, 6, 2},
		{1, 5, 3, 0},
	}

	for _, c := range cases {
		got := findClosest(c.x, c.y, c.z)
		if got != c.expected {

			t.Errorf("findClosest(%d, %d, %d) != %d, got %d\n", c.x, c.y, c.z, c.expected, got)
		}
	}
}

func TestConvertDateToBinary(t *testing.T) {
	cases := []struct {
		input    string
		expected string
	}{
		{"2080-02-29", "100000100000-10-11101"},
		{"1900-01-01", "11101101100-1-1"},
	}

	for _, c := range cases {
		got := convertDateToBinary(c.input)
		if got != c.expected {
			t.Errorf("convertDateToBinary(%q) != %q, got %q\n", c.input, c.expected, got)
		}
	}
}
