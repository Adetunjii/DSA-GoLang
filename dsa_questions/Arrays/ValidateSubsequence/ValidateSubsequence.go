package dsaquestions

func IsValidSubsequence(array []int, sequence []int) bool {
	seqIndex := 0

	for _, value := range array {
		if seqIndex == len(sequence) {
			break
		}
		if value == sequence[seqIndex] {
			seqIndex += 1
		}
	}
	return seqIndex == len(sequence)
}
