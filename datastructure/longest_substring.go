package datastructure

func LongestSubstring() (string, int) {

	s := "abcdefghijktir"

	maxLength := 0
	maxSubstring := ""
	charMap := make(map[byte]int)
	left, right := 0, 0

	for right < len(s) {
		if index, exists := charMap[s[right]]; exists && index >= left {
			left = index + 1
		}
		charMap[s[right]] = right
		if right-left+1 > maxLength {
			maxLength = right - left + 1
			maxSubstring = s[left : right+1]
		}
		right++
	}

	return maxSubstring, maxLength
}
