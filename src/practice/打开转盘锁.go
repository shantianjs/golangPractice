package practice

import "strings"

func openLock(deadends []string, target string) int {
	visited := make(map[string]bool,0)
	//使用map检查
	deadMap := make(map[string]bool, len(deadends))
	for _, str := range deadends {
		deadMap[str] = true
	}
	//初始状态，双向遍历
	workMap := map[string]bool{"0000": true}
	targetMap := map[string]bool{target: true}
	step := 0
	for len(workMap) > 0 {
		tem := map[string]bool{}
		for str := range workMap {
			//将判断，和状态转换的放在一起，visited不可放在修改选择列表内
			if _, ok := targetMap[str];ok {
				return step
			}
			if _,ok := deadMap[str]; ok {
				continue
			}
			//搜索当前数字
			visited[str] = true
			for _, nextStr := range getNetStrings(str) {
				if _,ok := visited[nextStr]; ok {
					continue
				}
				tem[nextStr] = true
			}
		}
		workMap = targetMap
		targetMap = tem
		step++
	}
	return -1
}

func getNetStrings(str string)(res []string){
	//获取当前状态的下一状态集合
	up := func (idx int) string {
		strlist := strings.Split(str, "")
		//加
		if strlist[idx] == "9" {
			strlist[idx] = "0"
		} else {
			strlist[idx] = string(strlist[idx][0] + 1)
		}
		return strings.Join(strlist, "")
	}
	down := func (idx int) string {
		strlist := strings.Split(str, "")
		//减
		if strlist[idx] == "0" {
			strlist[idx] = "9"
		} else {
			strlist[idx] = string(strlist[idx][0] - 1)
		}
		return strings.Join(strlist, "")
	}
	//遍历转动
	for idx, _ := range str {
		res = append(res, up(idx))
		res = append(res, down(idx))
	}
	return
}
