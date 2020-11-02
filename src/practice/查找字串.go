package practice

func strStr(haystack string, needle string) int {
	if needle == "" {
		return 0
	} else if haystack == "" || len(needle) > len(haystack){
		return -1
	}
	nLen := len(needle)
	for i := 0; i < len(haystack) - nLen + 1; i++ {
		if needle == haystack[i:i + nLen] {
			return i
		}
	}
	return -1
}
