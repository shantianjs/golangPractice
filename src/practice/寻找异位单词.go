package practice

func findAnagrams(s string, p string) (res []int) {
	wordLen := len(p)
	//需要检验的字符对应个数
	record := map[uint8]int{}
	//原始记录
	record1 := map[uint8]int{}
	for _, chr := range p {
		record[uint8(chr)]++
		record1[uint8(chr)]++
	}
	left, right := 0, 0
	valid := 0
	for right < len(s) {
		//扩大右边
		if _, ok := record[s[right]]; ok {
			if record[s[right]] <= record1[s[right]] && record[s[right]] > 0 {
				valid++
			}
			record[s[right]]--
		}
		right++
		//缩小左边
		for right - left > wordLen {
			if _, ok := record[s[left]]; ok {
				//个数必须小于合法字符数,大于零
				if record[s[left]] <= record1[s[left]] && record[s[left]] >= 0 {
					valid--
				}
				record[s[left]]++
			}
			left++
		}
		if valid >= wordLen {
			res = append(res, left)
		}
	}
	return
}
