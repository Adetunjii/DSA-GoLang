package minimumWindowSubString

import (
	"fmt"
	"math"
)

// Brute Force Approach O(N)
func MinimumWindowSubString(searchString, t string) string {

	minimumWindow := ""
	currentWindowLength := math.MaxInt32

	for i := 0; i < len(searchString); i++ {
		for j := i; j < len(searchString); j++ {

			currentWindow := subString(searchString, i, j+1)
			windowLength := len(currentWindow)

			matchFound := compareSubstringToCurrentWindow(currentWindow, t)

			if matchFound && windowLength < currentWindowLength {
				currentWindowLength = windowLength
				minimumWindow = currentWindow
			}
		}
	}

	return minimumWindow
}

func compareSubstringToCurrentWindow(currentWindow, subString string) bool {
	subStringChars := []rune(subString)
	currentWindowChars := []rune(currentWindow)

	hashMap := map[rune]int{}

	for i := 0; i < len(subStringChars); i++ {
		hashMap[subStringChars[i]] += 1
	}
	fmt.Println(hashMap)
	for i := 0; i < len(currentWindowChars); i++ {
		if value, ok := hashMap[currentWindowChars[i]]; ok {

			if value-1 == 0 {
				delete(hashMap, currentWindowChars[i])
			} else {
				hashMap[currentWindowChars[i]] = value - 1
			}
		}
	}

	fmt.Println(hashMap)
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
