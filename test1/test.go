package main

import "fmt"

func main() {
	var nums = []int{1}
	fmt.Println(nums[:0])
	fmt.Println(nums[1:])
}