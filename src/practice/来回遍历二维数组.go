package practice

func findDiagonalOrder(matrix [][]int) []int {
	//+x -y
	x, y := 0, 0
	stepx, stepy := -1, 1
	nr := len(matrix)
	nc := len(matrix[0])
	res := make([]int, 0)
	for x < nr && y < nc {
		res = append(res, matrix[x][y])
		if x+stepx < 0 || x+stepx >= nr {
			if y < nc - 1 {
				y += 1
			} else {
				x += 1
			}
			stepx, stepy = stepy, stepx
		} else if y+stepy < 0 || y+stepy >= nc {
			if x < nr - 1 {
				x += 1
			} else {
				y += 1
			}
			stepx, stepy = stepy, stepx
		} else {
			x += stepx
			y += stepy
		}
	}
	return res
}
