package practice

func checkInclusion(s1 string, s2 string) bool {
	origRecord := map[uint8]int{}
	record := map[uint8]int{}
	for _, chr := range s1 {
		origRecord[uint8(chr)]++
		record[uint8(chr)]++
	}
	valid := 0
	//左闭右开区间
	left, right := 0, 0
	for right < len(s2) {
		//扩张右侧s1的字符，因为是开区间
		if _, ok := record[s2[right]]; ok {
			//在S1中的都做记录，但是只有合法的才加检验次数，最小是1到0
			if record[s2[right]] > 0 {
				valid++
			}
			record[s2[right]]--
		}
		right++
		//大于子串长度，收缩左侧
		for right-left > len(s1) {
			//只有在子串中有当前元素。判断是否应该减少检验值
			if _, ok := record[s2[left]]; ok && record[s2[left]] < origRecord[s2[left]] {
				//从0到1个合法数字也是合法的
				if record[s2[left]] >= 0 {
					valid--
				}
				record[s2[left]]++
			}
			left++
		}
		//认证字符串长度等于子串，合法
		if valid == len(s1) {
			return true
		}
	}
	return false
}
