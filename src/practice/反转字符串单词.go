package practice

import "strings"

func reverseWords(s string) string {
	splitStr := strings.Split(s, " ")
	filteredStr := []string{}
	for _, str := range splitStr {
		if str != "" {
			filteredStr = append(filteredStr, str)
		}
	}
	n := len(filteredStr)
	for i := 0; i < n/2 ; i++ {
		filteredStr[i], filteredStr[n - i -1] = filteredStr[n - i -1], filteredStr[i]
	}
	s = strings.Join(filteredStr, " ")
	return s
}
