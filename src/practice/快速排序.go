package practice

func quickSort(arr []int) []int {
	var sort func(left, right int)
	sort = func(left, right int) {
		if left < right {
			pivot := partition(&arr, left, right)
			sort(left, pivot)
			sort(pivot + 1, right)
		}
	}
	sort(0 , len(arr) - 1)
	return arr
}

func partition(arr *[]int, left, right int) int {
	pivot := left
	index := left + 1
	for i := index; i <= right;i++ {
		if (*arr)[i] < (*arr)[pivot] {
			(*arr)[i], (*arr)[index] = (*arr)[index], (*arr)[i]
			index++
		}
	}
	(*arr)[pivot], (*arr)[index - 1] = (*arr)[index - 1], (*arr)[pivot]
	return index - 1
}
