package practice

import "unicode"

func IsPalindrome(s string) (bool, string) {
	var letters []rune
	for _, r := range s {
		if unicode.IsLetter(r) {
			letters = append(letters, unicode.ToLower(r))
		}
	}
	if len(letters) == 0 {
		return false,""
	}
	for i := range letters {
		if letters[i] != letters[len(letters)-1-i] {
			return false, string(letters)
		}
	}
	return true,string(letters)
}