package leetcode

import (
	"fmt"
	"strconv"
)

// Shared Helpers

func countChars(s string) map[rune]int {
	charMap := make(map[rune]int)
	for ch := 'a'; ch <= 'z'; ch++ {
		charMap[ch] = 0
	}

	for _, ch := range s {
		charMap[ch]++
	}

	return charMap
}

func contains(s string, ch rune) bool {
	for _, c := range s {
		if c == ch {
			return true
		}
	}
	return false
}

func toBinary(n int) string {
	if n == 0 {
		return "0"
	}

	result := ""
	for n > 0 {
		bit := n % 2
		result = strconv.Itoa(bit) + result
		n /= 2
	}
	return result
}

func intToDigits(n, numDigits int) []int {
	digits := make([]int, numDigits)

	divisor := 10
	finalIndex := numDigits - 1
	for i := 0; i < numDigits; i++ {
		rem := n % 10
		digits[finalIndex-i] = rem

		if n < 10 {
			break
		}

		n /= divisor
	}

	return digits
}

func countDigits(n int) []int {
	digits := intToDigits(n, 10)
	counts := make([]int, 10)

	inLeadingZeroes := true
	for i, digit := range digits {
		if inLeadingZeroes {
			if digit == 0 {
				if 9 == i {
					counts[0] = 1
				}
			} else {
				inLeadingZeroes = false
			}
		}

		if !inLeadingZeroes {
			counts[digit]++
		}
	}

	return counts
}

// https://leetcode.com/problems/convert-the-temperature/

func convertTemperature(celsius float64) []float64 {
	return []float64{celsius + 273.15, celsius*1.8 + 32}
}

// https://leetcode.com/problems/find-most-frequent-vowel-and-consonant/

func maxFreqSum(s string) int {
	charMap := countChars(s)

	maxVowel := 0
	vowels := "aeiou"
	for _, ch := range vowels {
		if charMap[ch] > maxVowel {
			maxVowel = charMap[ch]
		}
	}

	maxConsonant := 0
	for ch, count := range charMap {
		if !contains(vowels, ch) && count > maxConsonant {
			maxConsonant = count
		}
	}

	return maxVowel + maxConsonant
}

// https://leetcode.com/problems/filter-characters-by-frequency/

func filterCharacters(s string, k int) string {
	charMap := countChars(s)
	var result []rune
	for _, ch := range s {
		if charMap[ch] < k {
			result = append(result, ch)
		}
	}
	return string(result)
}

// https://leetcode.com/problems/compute-alternating-sum/

func alternatingSum(nums []int) int {
	sum := 0

	for i, v := range nums {
		if i%2 == 0 {
			sum += v
		} else {
			sum -= v
		}
	}

	return sum
}

// https://leetcode.com/problems/find-closest-person/

func findClosest(x int, y int, z int) int {
	diffX := z - x

	if diffX < 0 {
		diffX = -diffX
	}

	diffY := z - y

	if diffY < 0 {
		diffY = -diffY
	}

	if diffX < diffY {
		return 1
	}

	if diffY < diffX {
		return 2
	}

	return 0
}

// https://leetcode.com/problems/convert-date-to-binary/

func convertDateToBinary(date string) string {
	year := date[0:4]
	month := date[5:7]
	day := date[8:10]

	yearInt, _ := strconv.Atoi(year)
	monthInt, _ := strconv.Atoi(month)
	dayInt, _ := strconv.Atoi(day)

	yearBin := toBinary(yearInt)
	monthBin := toBinary(monthInt)
	dayBin := toBinary(dayInt)

	return fmt.Sprintf("%s-%s-%s", yearBin, monthBin, dayBin)
}

// https://leetcode.com/problems/reverse-degree-of-a-string/

func reverseDegree(s string) int {
	sum := 0

	for i, r := range s {
		diffA := r - 'a'
		diffZ := 26 - int(diffA)
		sum += (1 + i) * diffZ
	}

	return sum
}

// https://leetcode.com/problems/count-the-number-of-consistent-strings/

func countConsistentStrings(allowed string, words []string) int {
	count := 0

	for _, word := range words {
		notAllowedFound := false

		for i := 0; i < len(word) && !notAllowedFound; i++ {
			if !contains(allowed, rune(word[i])) {
				notAllowedFound = true
			}
		}

		if !notAllowedFound {
			count++
		}
	}

	return count
}

// https://leetcode.com/problems/smallest-even-multiple/

func smallestEvenMultiple(n int) int {
	if n%2 == 0 {
		return n
	}
	return n * 2
}

// https://leetcode.com/problems/find-the-degree-of-each-vertex/

func findDegrees(matrix [][]int) []int {
	degrees := make([]int, len(matrix))

	for i := range matrix {
		degree := 0
		for j := range matrix[i] {
			degree += matrix[i][j]
		}
		degrees[i] = degree
	}
	return degrees
}

// https://leetcode.com/problems/permutation-difference-between-two-strings/

func findPermutationDifference(s string, t string) int {
	total := 0

	for i := 0; i < len(s); i++ {
		for j := 0; j < len(t); j++ {
			if s[i] == t[j] {
				diff := i - j

				if diff > 0 {
					total += diff
				} else {
					total -= diff
				}
			}
		}
	}
	return total
}

// https://leetcode.com/problems/single-row-keyboard/

func calculateTime(keyboard string, word string) int {
	time := 0

	previous := 0
	for i := range word {
		for j := range keyboard {
			if word[i] == keyboard[j] {
				if j > previous {
					time += j - previous
				} else {
					time += previous - j
				}
				previous = j
				break
			}
		}
	}

	return time
}

// https://leetcode.com/problems/digit-frequency-score/

func digitFrequencyScore(n int) int {
	digitCounts := countDigits(n)

	score := 0

	for i, d := range digitCounts {
		score += i * d
	}

	return score
}

// https://leetcode.com/problems/minimum-element-after-replacement-with-digit-sum/

func minElement(nums []int) int {
	minEle := 46

	for _, num := range nums {
		digits := intToDigits(num, 5)

		sum := 0

		for _, d := range digits {
			sum += d
		}

		if sum < minEle {
			minEle = sum
		}
	}

	return minEle
}
