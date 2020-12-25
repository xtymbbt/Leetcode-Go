package main

import (
	"fmt"
	"math"
	"strconv"
)

// TODO: Problem code.
func largestPalindrome(n int) int {
	max := 0
	num := 0
	for i := int(math.Pow10(n)) - 1 ;i> int(math.Pow10(n-1))-1;i--{
		for j:=i;j>int(math.Pow10(n-1))-1;j--{
			a := i*j
			if check(a) && a > max {
				max = a
				num++
				if num > 2 {
					return max % 1337
				}
			}
		}
	}
	return max % 1337
}

func check(a int) bool {
	//fmt.Println("a is: ", a)
	str := strconv.Itoa(a)
	for i := 0; i < len(str)/2; i++ {
		if str[i] == str[len(str) - 1 - i] {
			continue
		} else {
			return false
		}
	}
	return true
}

func main() {
	//fmt.Println(check(906609))
	fmt.Println(largestPalindrome(6))
}