package isPalindrome

import (
	"testing"
)

func TestIsPalindrome(t *testing.T) {
	cases := []struct {
		x        int
		expected bool
	}{
		{-121, false},
		{121, true},
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

func TestGetDigits(t *testing.T) {
	
}
