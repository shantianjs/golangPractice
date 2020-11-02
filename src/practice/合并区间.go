package practice

import "sort"

func merge(intervals [][]int) [][]int {
	//排序
	sort.Slice(intervals, func(i, j int)bool {
		if intervals[i][0] != intervals[j][0] {
			return intervals[i][0] < intervals[j][0]
		}
		return intervals[i][1] < intervals[j][1]
	})
	//合并
	res := [][]int{}
	for _, item := range intervals {
		if len(res) == 0 {
			res = append(res, item)
			continue
		}
		left := res[len(res)-1]
		if left[1] >= item[0] {
			if item[1] > left[1] {
				left[1] = item[1]
			}
		} else {
			res = append(res, item)
		}
	}
	return res
}
