package _95

var res = 0
var present = 0

func maxAreaOfIsland(grid [][]int) int {
	res = 0
	present = 0
	for line := range grid {
		for column := range grid[line] {
			if grid[line][column] == 0 {
				continue
			}
			dfs(grid, line, column)
			if present > res {
				res = present
			}
			present = 0
		}
	}
	return res
}

func dfs(grid [][]int, line, column int) {
	if line < 0 || line >= len(grid) || column < 0 || column >= len(grid[0]) {
		return
	}
	if grid[line][column] == 0 {
		return
	}
	present++
	grid[line][column] = 0
	dfs(grid, line+1, column)
	dfs(grid, line-1, column)
	dfs(grid, line, column+1)
	dfs(grid, line, column-1)
}