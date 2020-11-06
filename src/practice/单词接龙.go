package practice

type Word struct {
	word string
	transNum int
}

func ladderLength(beginWord string, endWord string, wordList []string) int {
	has := false
	for _, word := range wordList {
		if word == endWord {
			has = true
			break
		}
	}
	if !has {
		return 0
	}

	search := func(word string) []string{
		res := []string{}
		newWordList := []string{}
		for idx, _ := range wordList {
			num := 0
			for i, _ := range word {
				if word[i] != wordList[idx][i] {
					num++
					if num > 1 {
						break
					}
				}
			}
			if num == 1 {
				res = append(res, wordList[idx])
			} else {
				newWordList = append(newWordList, wordList[idx])
			}
		}
		wordList = newWordList
		return res
	}
	workList := []Word{{beginWord,1}}
	for idx:= 0; idx < len(workList); idx++ {
		for _, item := range search(workList[idx].word) {
			if item == endWord {
				return workList[idx].transNum + 1
			}
			workList = append(workList, Word{item, workList[idx].transNum + 1})
		}
	}
	return 0
}
