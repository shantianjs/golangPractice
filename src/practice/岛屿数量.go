package practice

var(
	nr, nc int
)

func numIslands(grid [][]byte) int {
	res := 0
	nr = len(grid)
	nc = len(grid[0])
	for x:=0; x < nr; x++{
		for y := 0; y < nc; y++ {
			if grid[x][y] == '1' {
				clearOneLand(x, y, grid)
				res++
			}
		}
	}
	return res
}

//清除坐标周围的陆地
func clearOneLand(i, j int, b [][]byte) {
	if i < 0 || j < 0 || j >= nc || i >= nr || b[i][j] != '1' {
		return
	}
	b[i][j] = '0'
	clearOneLand(i, j + 1, b)
	clearOneLand(i, j - 1, b)
	clearOneLand(i - 1, j , b)
	clearOneLand(i + 1, j , b)
}

