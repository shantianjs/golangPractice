package practice

func validMountainArray(A []int) bool {
	if len(A) < 3 {
		return false
	}
	upFlag := [2]bool{false, false}
	for idx,num := range A {
		if idx == 0 {
			continue
		}
		if A[idx - 1] < num && !upFlag[1] {
			upFlag[0] = true
			continue
		}
		if A[idx - 1] > num &&  upFlag[0]{
			upFlag[1] = true
			continue
		}
		if A[idx -1] > num && upFlag[1] && upFlag[0]{
			continue
		}
		return false
	}
	return upFlag[0] && upFlag[1]
}
