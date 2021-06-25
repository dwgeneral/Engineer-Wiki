package main

import "fmt"

func bubbleSort(nums []int) {
	swapped := true
	for i := 0; i < len(nums); i++ {
		if !swapped {
			break
		}
		swapped = false
		for j := 1; j < len(nums)-i; j++ {
			if nums[j] < nums[j-1] {
				nums[j], nums[j-1] = nums[j-1], nums[j]
				swapped = true
			}
		}
	}
}

func main() {
	nums := []int{1, 5, 3, 2, 4}
	bubbleSort(nums)
	fmt.Println(nums)
}
