package practice

import (
	"strings"
)

func wordBreak(s string, wordDict []string) (sentences []string) {
	dp := make([][]string, len(s))
	n := len(s)

	var searchWrod func(s string, idx int) ([]string)
	searchWrod = func(s string, idx int) ([]string) {
		//dp回溯
		if dp[idx] != nil {
			return dp[idx][:]
		}
		result := []string{}
		//检查当前向后可匹配单词
		for _, word := range checkWords(s, wordDict) {
			//最后一个单词返回,不再向后查找
			if idx + len(word) == n {
				result = append(result, word)
				continue
			}
			//继续查找加入结果
			for _, suffix := range searchWrod(s[len(word):], idx + len(word))  {
				result = append(result, word + " " + suffix)
			}
		}
		// 记录结果
		dp[idx] = result
		return result
	}
	searchWrod(s, 0)
	return dp[0]
}

//检查后面是否可以继续
func checkWords(s string, wordDict []string) (res []string) {
	for _, word := range wordDict {
		if strings.HasPrefix(s, word) {
			res = append(res, word)
		}
	}
	return
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

