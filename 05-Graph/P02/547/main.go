package main

func main() {
	
}

func findCircleNum(M [][]int) int {
	ans, v := 0, make([]bool, len(M))
	var f func(int); f = func(i int) {
		for j := range M {
			if M[i][j] == 1 && !v[j] {
				v[j] = true
				f(j)
			}
		}
	}
	for i := range M {
		if !v[i] {
			v[i] = true
			f(i)
			ans++
		}
	}
	return ans
}
