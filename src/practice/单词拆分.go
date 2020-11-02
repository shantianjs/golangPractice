package practice

import "strings"

func wordBreak(s string, wordDict []string) []string {
	wordsMap := map[string]bool{}
	var result []string
	for _, word := range wordDict {
		wordsMap[word] = true
	}
	searchWrod(s, wordsMap, "", &result)
	return result
}

func searchWrod(s string, wordDict map[string]bool, preStr string, result *[]string) {
	if s == "" {
		*result = append(*result, preStr)
	}
	preWords := checkWords(s, wordDict)
	//回溯
	if len(preWords) == 0 {
		return
	}
	for _, word := range preWords {
		if preStr != "" {
			searchWrod(s[len(word):], wordDict, preStr + " " + word, result)
		} else {
			searchWrod(s[len(word):], wordDict,  word, result)
		}

	}
}

//检查后面是否可以继续
func checkWords(s string, wordDict map[string]bool) []string {
	res := []string{}
	for word := range wordDict {
		if strings.HasPrefix(s, word) {
			res = append(res, word)
		}
	}
	return res
}
//
//func wordBreak(s string, wordDict []string) (sentences []string) {
//	wordSet := map[string]struct{}{}
//	for _, w := range wordDict {
//		wordSet[w] = struct{}{}
//	}
//
//	n := len(s)
//	dp := make([][][]string, n)
//	var backtrack func(index int) [][]string
//	backtrack = func(index int) [][]string {
//		if dp[index] != nil {
//			return dp[index]
//		}
//		wordsList := [][]string{}
//		for i := index + 1; i < n; i++ {
//			word := s[index:i]
//			if _, has := wordSet[word]; has {
//				for _, nextWords := range backtrack(i) {
//					wordsList = append(wordsList, append([]string{word}, nextWords...))
//				}
//			}
//		}
//		word := s[index:]
//		if _, has := wordSet[word]; has {
//			wordsList = append(wordsList, []string{word})
//		}
//		dp[index] = wordsList
//		return wordsList
//	}
//	for _, words := range backtrack(0) {
//		sentences = append(sentences, strings.Join(words, " "))
//	}
//	return
//}

