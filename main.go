package main

import (
	"fmt"

	"github.com/seyf97/leetCode/solutions"
)

var print = fmt.Println

func main() {
	nums := []int{1, 1, 2, 2, 3, 3, 3, 4, 5, 5}
	solutions.RemoveElement(nums, 1)
	print(nums)
}
