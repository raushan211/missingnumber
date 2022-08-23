package main

import "fmt"

func main() {
	nums := []int{3, 0, 1}
	arr := bubblesort(nums)

	fmt.Println(missingNumber(arr))

}
func missingNumber(nums []int) int {
	for i := 0; i < len(nums); i++ {
		if i != nums[i] {
			return i
		}
	}
	return -1
}

func bubblesort(nums []int) []int {

	for i := 0; i < len(nums); i++ {

		for j := 0; j < len(nums)-1; j++ {

			if nums[j] > nums[j+1] {

				nums = swap(nums, j, j+1)
			}

		}
	}
	return nums
}

func swap(nums []int, val1, val2 int) []int {

	temp := nums[val1]

	nums[val1] = nums[val2]
	nums[val2] = temp

	return nums

}
