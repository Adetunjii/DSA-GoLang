package plusOne

func PlusOne(digits []int) []int {

	right := len(digits) - 1
	carry := 1

	for i := right; i >= 0; i-- {
		sum := digits[i] + carry
		digits[i] = sum % 10
		carry = sum / 10
	}

	if carry > 0 {
		return append([]int{carry}, digits...)
	}
	return digits
}
