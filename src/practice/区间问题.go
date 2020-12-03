package practice

import "sort"

func removeCoveredIntervals(intervals [][]int) int {
	if len(intervals) == 0 {
		return 0
	}
	sort.Slice(intervals, func(i, j int) bool {
		a, b := intervals[i], intervals[j]
		if a[0] != b[0] {
			return a[0] < b[0]
		}
		return a[1]>b[1]
	})
	pre := intervals[0]
	count := 1
	for i := 1; i < len(intervals); i++ {
		cur := intervals[i]
		if pre[0] <= cur[0] && pre[1] >= cur[1] {
			continue
		}
		pre = cur
		count++
	}
	return count
}

func merge(intervals [][]int) [][]int {
	if len(intervals) == 0 {
		return nil
	}
	sort.Slice(intervals, func(i, j int) bool {
		a, b := intervals[i], intervals[j]
		if a[0] != b[0] {
			return a[0] < b[0]
		}
		return a[1]>b[1]
	})
	pre := intervals[0]
	res := [][]int{intervals[0]}
	for i := 1; i < len(intervals); i++ {
		cur := intervals[i]

		//不相交的
		if cur[0] > pre[1] {
			res = append(res, cur)
			pre = cur
		} else {
			//排序后前面的必定0更小
			res[len(res)-1][1] = max(cur[1], pre[1])
		}
	}
	return res
}

func intervalIntersection(A [][]int, B [][]int) [][]int {
	if len(A) == 0 {
		return nil
	}
	min := func(x, y int) int {
		if x < y {
			return x
		}
		return y
	}
	max := func(x, y int) int {
		if x > y {
			return x
		}
		return y
	}
	res := [][]int{}
	i, j := 0, 0
	for i <= len(A) && j <= len(B) {
		//判断是否有交集
		if A[i][1] > B[j][0] || B[j][1] > A[i][0] {
			a, b := max(A[i][0], B[j][0]), min(A[i][1], B[j][1])
			if a <= b {
				res = append(res, []int{a,b})
			}
		}
		//移动指针
		if A[i][1] < B[j][1] {
			i++
		} else {
			j++
		}
	}
	return res
}