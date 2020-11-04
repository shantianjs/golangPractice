package practice

func MergeSort(arr []int) []int{
	if len(arr) <= 1 {
		return arr
	}
	idx := len(arr)/2
	return _merge(MergeSort(arr[:idx]), MergeSort(arr[idx:]))
}

func _merge(left , right []int) []int {
	res := make([]int, len(left) + len(right))
	var idx, idy int
	for idz := 0; idz < len(res) ; idz++ {
		if idx < len(left) && idy < len(right) {
			if left[idx] < right[idy] {
				res[idz] = left[idx]
				idx++
			} else {
				res[idz] = right[idy]
				idy++
			}
			continue
		}
		if idx >= len(left) {
			copy(res[idz:], right[idy:])
		} else {
			copy(res[idz:], left[idx:])
		}
		break
	}
	return res
}
