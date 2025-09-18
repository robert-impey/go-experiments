package leetcode

// https://leetcode.com/problems/convert-the-temperature/

func convertTemperature(celsius float64) []float64 {
	return []float64{celsius + 273.15, celsius*1.8 + 32}
}

// https://leetcode.com/problems/find-most-frequent-vowel-and-consonant/

func maxFreqSum(s string) int {
	charMap := make(map[rune]int)
	for ch := 'a'; ch <= 'z'; ch++ {
		charMap[ch] = 0
	}

	for _, ch := range s {
		charMap[ch]++
	}

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

func contains(s string, ch rune) bool {
	for _, c := range s {
		if c == ch {
			return true
		}
	}
	return false
}
