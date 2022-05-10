package ValidAnagram

func ValidAnagram(s, t string) bool {
	if len(s) != len(t) {
		return false
	}

	hashMap := make(map[byte]int)

	for i := 0; i < len(s); i++ {

		hashMap[s[i]] += 1
	}

	newT := []byte(t)
	for j := 0; j < len(newT); j++ {
		if value, ok := hashMap[t[j]]; ok {

			if value-1 == 0 {
				delete(hashMap, t[j])
			} else {
				value -= 1
				hashMap[t[j]] = value
			}
		}
	}

	return len(hashMap) == 0
}
