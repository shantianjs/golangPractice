package practice

func isValid(s string) bool {
	record := map[int32]int32{'(': ')', '[': ']', '{': '}'}
	stack := []int32{}
	for _, chr := range s {
		if len(stack) == 0 {
			if _, ok := record[chr]; ok {
				stack = append(stack, chr)
			} else {
				return false
			}
		} else {
			if _, ok := record[chr]; ok {
				stack = append(stack, chr)
			} else {
				if chr == record[stack[len(stack)-1]] {
					stack = stack[:len(stack)-1]
				} else {
					return false
				}

			}
		}
	}
	return true
}
