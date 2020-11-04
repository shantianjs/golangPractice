package practice

func insert(intervals [][]int, newInterval []int) [][]int {
	if len(intervals) == 0 {
		intervals = append(intervals, newInterval)
		return intervals
	}
	//插入
	var merge func(idx int)
	merge = func(idx int) {
		minIdx := idx
		maxIdx := idx
		for idy := idx; idy >= 0; idy-- {
			if intervals[idy][1] >= newInterval[0] {
				minIdx = idy
			}
		}
		for idy := idx; idy < len(intervals); idy++ {
			if intervals[idy][0] <= newInterval[1]{
				maxIdx = idy
			}
		}
		if minIdx != maxIdx {
			intervals[minIdx][1] = - min(-intervals[maxIdx][1] , -intervals[minIdx][1], -newInterval[1])
			intervals[minIdx][0] = min(intervals[maxIdx][0], intervals[minIdx][0], newInterval[0])
			intervals = append(intervals[:minIdx+1], intervals[maxIdx+1:]...)
		}

	}
	for idx, item := range intervals {
		if item[0] > newInterval[0] {
			intervals = append(intervals, []int{})
			copy(intervals[idx+1:], intervals[idx:])
			intervals[idx] = newInterval
			merge(idx)
			break
		}
		if idx == len(intervals) - 1 {
			intervals = append(intervals, newInterval)
			merge(idx+1)
		}
	}

	return intervals
}

func min(x ...int) int {
	i := x[0]
	for _, j := range x {
		if j < i {
			i = j
		}
	}
	return i
}
