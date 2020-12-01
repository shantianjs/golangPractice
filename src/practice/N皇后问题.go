package practice

import (
	"math"
	"strings"
)

var res [][]string

func solveNQueens(n int) [][]string {
	//初始化
	res = make([][]string,0)
	place(0,make([]string, n, n))
	return res
}

//开始搜索对应行
func place(row int,array []string) {
	//最后一个结束
	if row == len(array) {
		temp := make([]string, len(array))
		copy(temp, array)
		res = append(res, temp)
		return
	}
	n := len(array) - 1
	for n >= 0 {
		if queIsValid(row, n, array) {
			temp := array[:]
			temp[row] = generateStr(len(array), n)
			place(row + 1, temp)
		}
		n--
	}
}

func generateStr(maxLen,n int) string {
	sb := strings.Builder{}
	for i := 0; i < n ; i++ {
		sb.WriteString(".")
	}
	sb.WriteString("Q")
	for i := sb.Len(); i < maxLen; i++ {
		sb.WriteString(".")
	}
	return sb.String()
}

func queIsValid(row, col int, array []string) bool {
	//col 列,最新一行
	for s_row := 0; s_row < row; s_row++ {
		str := array[s_row]
		s_col := strings.Index(str, "Q")
		//左上,右上
		if int(math.Abs(float64(col - s_col))) == row - s_row {
			return false
		}
		//同列
		if s_col == col {
			return false
		}
	}
	return true
}
