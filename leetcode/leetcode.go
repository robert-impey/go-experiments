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
