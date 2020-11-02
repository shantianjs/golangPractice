package practice

func LongestPalindrome(s string) string {
	maxPalindrome := ""
	maxLen := 0
	for idx, _ := range s {
		i := idx - maxLen
		for i >= 0 {
			if idx < len(s) - 1 {
				if ok := isPalindrome(s[i:idx+1]);ok {
					maxLen = idx + 1 - i
					maxPalindrome = s[i:idx+1]
				}
			} else{
				if ok := isPalindrome(s[i:]);ok {
					maxLen = idx + 1 - i
					maxPalindrome = s[i:]
				}
			}
			i--
		}
	}
	return maxPalindrome
}

func isPalindrome(s string) bool {
	if len(s) == 0 {
		return false
	}
	for i, _ := range s {
		if i >= len(s) / 2 {
			break
		}
		if s[i] != s[len(s)-1-i] {
			return false
		}
	}
	return true
}