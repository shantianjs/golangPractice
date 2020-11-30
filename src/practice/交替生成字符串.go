package practice

import (
	"sort"
	"strings"
)

func reorganizeString(S string) string {
	//记录字符出现个数
	record := map[int32]int{}
	for _,chr := range S {
		record[chr]++
	}
	//计入字符合集
	keys := make([]int32, len(record), len(record))
	idx := 0
	for key := range record {
		keys[idx] = key
		idx++
	}
	//使字符递减排序
	sort.Slice(keys, func(i, j int) bool {
		return record[keys[i]] > record[keys[j]]
	})
	//判断最大个数字符
	if record[keys[0]] > (len(S) + 1)/2 {
		return ""
	}
	sb := make([]string, len(S), len(S))
	//keys的位置
	sIdx := 0
	fill := func(idx int) {
		for idx < len(S)  {
			sb[idx] = string(keys[sIdx])
			record[keys[sIdx]]--
			if record[keys[sIdx]] == 0 {
				sIdx++
			}
			idx += 2
		}
	}
	//先填充奇数位，当S是奇数个时奇数位是1/2+1，一定要让字符数目最大的先填充
	fill(0)
	//填充偶数数位
	fill(1)
	return strings.Join(sb,"")
}
