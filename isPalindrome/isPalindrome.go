package isPalindrome

import "math"

func IsPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	if x < 10 {
		return true
	}

	numDigits := countDigits(x)

	halfDigits := numDigits / 2

	for i := 0; i < halfDigits; i++ {
		j := numDigits - i - 1
		if getDigit(x, i) != getDigit(x, j) {
			return false
		}
	}

	return true
}

func intAbs(x int) int {
	if x < 0 {
		return -x
	} else {
		return x
	}
}

func getDigit(x int, i int) int {
	absX := intAbs(x)

	if i > countDigits(absX)-1 {
		return -1
	}

	tenToI := int(math.Pow10(i))
	xDivTenToI := absX / tenToI

	return xDivTenToI % 10
}

func countDigits(x int) int {
	absX := intAbs(x)
	numDigits := 1

	for ; ; numDigits++ {
		absX /= 10
		if absX == 0 {
			break
		}
	}

	return numDigits
}
