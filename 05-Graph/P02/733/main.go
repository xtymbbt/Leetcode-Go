package main

func main() {

}

func floodFill(image [][]int, sr int, sc int, newColor int) [][]int {
	target := image[sr][sc]
	if target == newColor {
		return image
	}
	dfs(&image, sr, sc, target, newColor)
	return image
}

func dfs(image *[][]int, sr, sc, target, newColor int) {
	if sr < 0 || sr >= len(*image) || sc < 0 || sc >= len((*image)[0]) {
		return
	}
	if (*image)[sr][sc] != target {
		return
	}
	(*image)[sr][sc] = newColor
	dfs(image, sr+1, sc, target, newColor)
	dfs(image, sr-1, sc, target, newColor)
	dfs(image, sr, sc+1, target, newColor)
	dfs(image, sr, sc-1, target, newColor)
}
