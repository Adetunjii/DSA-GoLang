package groupAnagrams

import "sort"

// Time O(s * wlogw) Space => O(sw)
func GroupAnagrams(strs []string) [][]string {
	hashMap := make(map[string][]string)
	var result [][]string

	if len(strs) == 0 {
		return result
	}

	for i := 0; i < len(strs); i++ {
		sortedWord := sortWord(strs[i])
		hashMap[sortedWord] = append(hashMap[sortedWord], strs[i])
	}

	for _, value := range hashMap {
		result = append(result, value)
	}

	return result
}

func sortWord(word string) string {
	wordBytes := []byte(word)
	sort.Slice(wordBytes, func(i, j int) bool {
		return wordBytes[i] < wordBytes[j]
	})

	return string(wordBytes)
}
