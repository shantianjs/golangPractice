package practice

import (
	"fmt"
	"testing"
)

func TestSamllSort(t *testing.T) {
	inputs := []struct{
		nums []int
		res  []int
	} {
		{[]int{8, 1, 2, 2, 3},[]int{4,0,1,1,3}},
		{[]int{6,5,4,8},[]int{2,1,0,3}},
		{[]int{7,7,7,7},[]int{0,0,0,0}},
	}
	for _, item := range inputs {
		fmt.Println(item.nums)
		res := smallerNumbersThanCurrent(item.nums)
		if !compare(item.res, res) {
			t.Errorf("输出:%v 正确：%v \n",item.nums, item.res )
		}
	}
}

func compare(s1, s2 []int) bool {
	for idx, i := range s1 {
		if s2[idx] != i {
			return false
		}
	}
	return true
}