package main

import (
	"fmt"
	"math"
)

func main() {
	max, min := maxMin(1, 3, 2, 5, 7, 9)
	fmt.Print(max, min)
}

func add(a, b int) (sum int) {
	sum = a + b
	return
}

func maxMin(nums ...int) (int, int) {
	if len(nums) == 0 {
		panic("必须至少包含一个数字！")
	}

	max := math.MinInt
	min := math.MaxInt

	for _, num := range nums {
		if max < num {
			max = num
		}
		if min > num {
			min = num
		}
	}

	return max, min
}
