package practice

func islandPerimeter(grid [][]int) int {
	if len(grid) == 0 {
		return 0
	}
	nr := len(grid)
	nc := len(grid[0])
	count := 0
	for x, row := range grid {
		for y, _ := range row {
			//中间水域
			if grid[x][y] == 0 {
				continue
			}
			//上
			if x == 0 || grid[x-1][y] == 0 {
				count += 1
			}
			//下
			if x == nr-1 || grid[x+1][y] == 0 {
				count += 1
			}
			//左
			if y == 0 || grid[x][y-1] == 0 {
				count += 1
			}
			//右
			if y == nc-1 || grid[x][y+1] == 0 {
				count += 1
			}
		}
	}
	return count
}