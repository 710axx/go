package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Print("Hello Git!")
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
