package minimumWindowSubString

import (
	"math"
)

// Brute Force Approach O(N ^ 2)
//func MinimumWindowSubString(searchString, t string) string {
//
//	sl := len(t)
//	minimumWindow := ""
//	currentWindowLength := math.MaxInt32
//
//	for i := 0; i < len(searchString); i++ {
//		for j := i; j < len(searchString); j++ {
//
//			currentWindow := subString(searchString, i, sl)
//			windowLength := len(currentWindow)
//
//			matchFound := compareSubstringToCurrentWindow(currentWindow, t)
//
//			if matchFound && windowLength < currentWindowLength {
//				currentWindowLength = windowLength
//				minimumWindow = currentWindow
//			}
//		}
//	}
//
//	return minimumWindow
//}

// Optimal Approach
func MinimumWindowSubString(searchString, t string) string {

	minWindowLength := math.MaxInt32
	left := 0
	right := 0
	strLen := len(searchString)
	minWindow := ""

	for right < strLen {
		currentWindow := subString(searchString, left, right-left+1)
		currentWindowLength := len(currentWindow)

		if compareSubstringToCurrentWindow(currentWindow, t) {

			if currentWindowLength < minWindowLength {
				minWindowLength = currentWindowLength
				minWindow = currentWindow
			}
			left += 1
		} else {
			right += 1
		}
	}

	return minWindow
}

func compareSubstringToCurrentWindow(currentWindow, subString string) bool {
	subStringChars := []rune(subString)
	currentWindowChars := []rune(currentWindow)

	hashMap := map[rune]int{}

	for i := 0; i < len(subStringChars); i++ {
		hashMap[subStringChars[i]] += 1
	}

	for i := 0; i < len(currentWindowChars); i++ {
		if value, ok := hashMap[currentWindowChars[i]]; ok {

			if value-1 == 0 {
				delete(hashMap, currentWindowChars[i])
			} else {
				hashMap[currentWindowChars[i]] = value - 1
			}
		}
	}
	return len(hashMap) == 0
}

func subString(input string, start int, length int) string {
	runes := []rune(input)

	if start > len(runes) {
		return ""
	}

	if start+length > len(runes) {
		length = len(runes) - start
	}

	return string(runes[start : start+length])
}
