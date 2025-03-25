package isPalindrome

import (
	"testing"
)

func TestIsPalindrome(t *testing.T) {
	cases := []struct {
		x        int
		expected bool
	}{
		{121, true},
		{-121, false},
		{0, true},
		{122, false},
	}

	for _, c := range cases {
		got := IsPalindrome(c.x)
		if got != c.expected {
			t.Errorf("IsPalindrome(%d) != %t, got %t\n", c.x, c.expected, got)
		}
	}
}

func TestGetDigit(t *testing.T) {
	cases := []struct {
		x, i, expected int
	}{
		{-121, 0, 1},
		{121, 1, 2},
		{0, 0, 0},
		{122, 2, 1},
		{123, 3, -1},
		{-123, 3, -1},
	}

	for _, c := range cases {
		got := getDigit(c.x, c.i)
		if got != c.expected {
			t.Errorf("getDigit(%d, %d) != %d, got %d\n", c.x, c.i, c.expected, got)
		}
	}
}

func TestCountDigits(t *testing.T) {
	cases := []struct {
		x, expected int
	}{
		{-121, 3},
		{121, 3},
		{0, 1},
		{122, 3},
		{123, 3},
		{-123, 3},
		{12345678, 8},
	}

	for _, c := range cases {
		got := countDigits(c.x)
		if got != c.expected {
			t.Errorf("countDigits(%d) != %d, got %d\n", c.x, c.expected, got)
		}
	}
}
