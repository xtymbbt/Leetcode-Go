package main

func main() {

}

func numIslands(grid [][]byte) int {
	count := 0
	for line := range grid {
		for column := range grid[line] {
			if grid[line][column] == '0' {
				continue
			}
			count++
			dfs(grid, line, column)
		}
	}
	return count
}

func dfs(grid [][]byte, line int, column int) {
	if line < 0 || column < 0 || line >= len(grid) || column >= len(grid[0]){
		return
	}
	if grid[line][column] == '0' {
		return
	}
	grid[line][column] = '0'
	dfs(grid, line+1, column)
	dfs(grid, line-1, column)
	dfs(grid, line, column+1)
	dfs(grid, line, column-1)
	return
}